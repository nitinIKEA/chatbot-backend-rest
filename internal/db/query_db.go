package db

import (
	"encoding/json"
	"fmt"

	"github.com/godror/godror"
)

type OrderStatus struct {
	BUCode            string        `json:"BU_CODE"`
	BUType            string        `json:"BU_TYPE"`
	ExecutableOrderID string        `json:"EXECUTABLE_ORDER_ID"`
	TotalOrderLine    godror.Number `json:"TOTAL_ORDER_LINE"`
	OrdStatus         int64         `json:"ORD_STATUS"`
}

type PickingGroup struct {
	StoreDetails string `json:"STORE_DETAILS"`
}

type AvailableStock struct {
	BUCode     string        `json:"BU_CODE"`
	BUType     string        `json:"BU_TYPE"`
	ItemNo     string        `json:"ITEM_NO"`
	ItemType   string        `json:"ITEM_TYPE"`
	BUCodeSup  string        `json:"BU_CODE_SUP"`
	BUTypeSup  string        `json:"BU_TYPE_SUP"`
	ItemAvail  godror.Number `json:"ITEM_AVAIL"`
	UomCodeQty string        `json:"UOM_CODE_QTY"`
}

type BlockedStock struct {
	BUCode       string        `json:"BU_CODE"`
	BUType       string        `json:"BU_TYPE"`
	ItemNo       string        `json:"ITEM_NO"`
	ItemType     string        `json:"ITEM_TYPE"`
	BUCodeSup    string        `json:"BU_CODE_SUP"`
	BUTypeSup    string        `json:"BU_TYPE_SUP"`
	ItemQtyBlock godror.Number `json:"ITEM_QTY_BLOCK"`
	UomCodeQty   string        `json:"UOM_CODE_QTY"`
}

func (c DBConns) QueryOrderstatusManordref(ordidrefreq, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryOrderstatusManordref"], ordidrefreq)
	data, err := c.orderStatusQuery(qry, env)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c DBConns) QueryOrderstatusOrdidref(orgordref, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryOrderstatusOrdidref"], orgordref)
	data, err := c.orderStatusQuery(qry, env)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c DBConns) QueryOrderstatusOrdref(ordIdRefSales string, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryOrderstatusOrdref"], ordIdRefSales)
	data, err := c.orderStatusQuery(qry, env)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c DBConns) QueryOrderstatusWrkordref(ordidrefwork, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryOrderstatusWrkordref"], ordidrefwork)
	data, err := c.orderStatusQuery(qry, env)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c DBConns) QueryPickingroup(storenum, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryPickingroup"], storenum, storenum, storenum, storenum, storenum)
	rw, err := c[env].Query(qry)
	if err != nil {
		return nil, err
	}
	var pg PickingGroup
	var pgs []PickingGroup
	for rw.Next() {
		if err := rw.Scan(&pg.StoreDetails); err != nil {
			return nil, err
		}
		pgs = append(pgs, pg)
	}
	dt, err := prepareDataJson(pgs)
	if err != nil {
		return nil, err
	}
	return dt, nil
}

func (c DBConns) QueryStockavailabilityAvailable(bucode, butype, itemnum, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryStockavailabilityAvailable"], bucode, butype, itemnum, itemnum, bucode, butype, itemnum, bucode, butype, bucode, butype, itemnum, itemnum, itemnum)
	rw, err := c[env].Query(qry)
	if err != nil {
		return nil, err
	}
	var stk AvailableStock
	var stks []AvailableStock
	for rw.Next() {
		if err := rw.Scan(&stk.BUCode, &stk.BUType, &stk.ItemNo, &stk.ItemType, &stk.BUCodeSup, &stk.BUTypeSup, &stk.ItemAvail, &stk.UomCodeQty); err != nil {
			return nil, err
		}
		stks = append(stks, stk)
	}
	dt, err := prepareDataJson(stks)
	if err != nil {
		return nil, err
	}
	return dt, nil
}

func (c DBConns) QueryStockavailabilityBlocked(bucode, butype, itemnum, env string) ([]byte, error) {
	qry := fmt.Sprintf(Queries["QueryStockavailabilityBlocked"], bucode, butype, itemnum)
	rw, err := c[env].Query(qry)
	if err != nil {
		return nil, err
	}
	var stk BlockedStock
	var stks []BlockedStock
	for rw.Next() {
		if err := rw.Scan(&stk.BUCode, &stk.BUType, &stk.ItemNo, &stk.ItemType, &stk.BUCodeSup, &stk.BUTypeSup, &stk.ItemQtyBlock, &stk.UomCodeQty); err != nil {
			return nil, err
		}
		stks = append(stks, stk)
	}
	dt, err := prepareDataJson(stks)
	if err != nil {
		return nil, err
	}
	return dt, nil
}

func prepareDataJson(data interface{}) ([]byte, error) {
	dt, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dt, nil
}

func (c DBConns) orderStatusQuery(qry, env string) ([]byte, error) {
	dest := OrderStatus{}
	data := []OrderStatus{}
	rows, err := c[env].Query(qry)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&dest.BUCode, &dest.BUType, &dest.ExecutableOrderID, &dest.TotalOrderLine, &dest.OrdStatus); err != nil {
			return nil, err
		}
		data = append(data, dest)
	}

	res, err := prepareDataJson(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
