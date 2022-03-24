package db

var Queries = map[string]string{
	"QueryOrderstatusManordref": `select lu.bu_code,lu.bu_type,o.ord_id executable_order_id,count(ol.ordl_id) total_order_line,o.ord_status
	from orders o
	join order_lines ol on (o.ord_id=ol.ord_id)
	join logistic_units lu on (o.lu_tk=lu.lu_tk)
	where o.ord_id_ref_req = '%s'
	group by lu.bu_code,lu.bu_type,o.ord_id,o.ord_status`,

	"QueryOrderstatusOrdidref": `select lu.bu_code,lu.bu_type,o.ord_id executable_order_id,count(ol.ordl_id) total_order_line,o.ord_status
	from orders o
	join order_lines ol on (o.ord_id=ol.ord_id)
	join logistic_units lu on (o.lu_tk=lu.lu_tk)
	where o.org_ord_ref = '%s'
	group by lu.bu_code,lu.bu_type,o.ord_id,o.ord_status`,

	"QueryOrderstatusOrdref": `select lu.bu_code,lu.bu_type,o.ord_id executable_order_id,count(ol.ordl_id) total_order_line,o.ord_status
	from orders o
	join order_lines ol on (o.ord_id=ol.ord_id)
	join logistic_units lu on (o.lu_tk=lu.lu_tk)
	where o.ord_id_ref_sales = '%s'
	group by lu.bu_code,lu.bu_type,o.ord_id,o.ord_status`,

	"QueryOrderstatusWrkordref": `select lu.bu_code,lu.bu_type,o.ord_id executable_order_id,count(ol.ordl_id) total_order_line,o.ord_status
	from orders o
	join order_lines ol on (o.ord_id=ol.ord_id)
	join logistic_units lu on (o.lu_tk=lu.lu_tk)
	where o.ord_id_ref_wrk = '%s'
	group by lu.bu_code,lu.bu_type,o.ord_id,o.ord_status`,

	"QueryStockavailabilityAvailable": `WITH
	sv
	AS
	  (SELECT    lu_tk
				,item_no
				,item_type
				,bu_code_sup
				,bu_type_sup
				,SUM(item_qty - item_qty_ddc_alloc - item_qty_stop_tot) AS item_avail
				,uom_code_qty
	   FROM      (SELECT  lu_tk
					  ,item_no
					  ,item_type
					  ,bu_code_sup
					  ,bu_type_sup
					  ,wis_data.wis_data_utility.pieces_or_meters(item_qty, uom_code_qty) AS item_qty
					  ,wis_data.wis_data_utility.uom_code(uom_code_qty) AS uom_code_qty
					  ,CASE
						 WHEN order_line_key IS NOT NULL THEN
						   wis_data_utility.pieces_or_meters(item_qty, uom_code_qty)
						 ELSE
						   0
					   END AS item_qty_ddc_alloc
					  ,CASE WHEN stock_status_list IS NOT NULL THEN item_qty ELSE 0 END AS item_qty_stop_tot
			   FROM    wis_data.stock_objects
			   WHERE   lu_tk = logistic_units_utl.lu_tk('%s', '%s')
			   AND     item_no = '%s'
			   AND     item_type = 'ART'
			   AND     stock_status_list IS NULL) sv
	   WHERE     item_no = '%s' AND item_type = 'ART'
	   GROUP BY  lu_tk
				,item_no
				,item_type
				,bu_code_sup
				,bu_type_sup
				,uom_code_qty),
	sr
	AS
	  (SELECT    lu_tk
				 ,item_no
				 ,item_type
				 ,bu_code_sup
				 ,bu_type_sup
				 ,-1 * SUM(item_qty_resv) AS item_avail
				 ,uom_code_qty
		FROM      wis_data.stock_resvs_vw
		WHERE     lu_tk = logistic_units_utl.lu_tk('%s', '%s')
		AND       item_no = '%s'
		AND       item_type = 'ART'
		GROUP BY  lu_tk
				 ,item_no
				 ,item_type
				 ,bu_code_sup
				 ,bu_type_sup
				 ,uom_code_qty
				 ,ord_src)
	   ,
	sa
	AS
	  (SELECT    lu_tk
				,item_no
				,item_type
				,bu_code_sup
				,bu_type_sup
				,SUM(-1 * item_qty_alloc) AS item_avail
				,uom_code_qty
	   FROM      (SELECT  o.lu_tk
					,ol.item_no
					,ol.item_type
					,ol.bu_code_sup
					,ol.bu_type_sup
					,wis_data_utility.pieces_or_meters(ol.item_qty_alloc, ol.uom_code_qty) AS item_qty_alloc -- hotfix
					,wis_data_utility.uom_code(ol.uom_code_qty) AS uom_code_qty
					,ol.ordl_id_ref_pur
			 FROM    (SELECT  *
					  FROM    orders
					  WHERE   lu_tk = logistic_units_utl.lu_tk('%s', '%s')) o
					 JOIN
					 (SELECT  *
					  FROM    wis_data.order_line_allocs
					  WHERE   lu_tk = logistic_units_utl.lu_tk('%s', '%s')
					  AND     item_no = '%s'
					  AND     item_type = 'ART') ol
					   ON o.ord_id = ol.ord_id)
	   WHERE     ordl_id_ref_pur IS NULL
	   AND       item_no = '%s'
	   AND       item_type = 'ART'
	   GROUP BY  lu_tk
				,item_no
				,item_type
				,bu_code_sup
				,bu_type_sup
				,uom_code_qty)
		SELECT    lu.bu_code,lu.bu_type,s.item_no
                                     ,s.item_type
                                     ,s.bu_code_sup
                                     ,s.bu_type_sup
                                     ,LEAST(SUM(s.item_avail), 999999999.99) AS item_avail
                                     ,s.uom_code_qty
                            FROM      (SELECT * FROM sv
                                       UNION ALL
                                       SELECT * FROM sa
                                       UNION ALL
                                       SELECT * FROM sr) s
                                       join logistic_units lu on (s.lu_tk=lu.lu_tk)
                            WHERE     s.item_no = '%s' AND s.item_type = 'ART'
                            GROUP BY  lu.bu_code,lu.bu_type,s.item_no
                                     ,s.item_type
                                     ,s.bu_code_sup
                                     ,s.bu_type_sup
                                     ,s.uom_code_qty`,

	"QueryStockavailabilityBlocked": `SELECT  lu.bu_code,lu.bu_type
		,item_no
		,item_type
		,bu_code_sup
		,bu_type_sup
		,LEAST(SUM(item_qty), 999999999.99) AS item_qty_block
		,wis_data.wis_data_utility.uom_code(uom_code_qty) AS uom_code_qty
		FROM    wis_data.stock_objects s
		join logistic_units lu on (s.lu_tk=lu.lu_tk)
		WHERE   s.lu_tk = logistic_units_utl.lu_tk('%s', '%s')
		AND item_no = '%s'
		AND     item_type = 'ART'
		AND     stock_status_list IS NOT NULL
		GROUP BY  lu.bu_code,lu.bu_type,item_no
		,item_type
		,bu_code_sup
		,bu_type_sup
		,uom_code_qty`,

	"QueryPickingroup": `select case 
		when  (select count(*) from logistic_units  where bu_code='%s' and bu_type='STO' and valid_to is null)>1 then '%s'|| ' Store has Configuaration issue'
		when  (select count(*) from logistic_units  where bu_code='%s' and bu_type='STO' and valid_to is null)=0 then '%s'||' Store not configured'
		else '%s'||' Store already configured'
		end Store_details from dual`,
}
