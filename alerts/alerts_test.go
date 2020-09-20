package alerts

import (
	"testing"
)

func TestErrors(t *testing.T) {
	ok, err := CheckAlert(">", "", "")
	if ok || err == nil {
		t.Fatalf("fail:%+v\n", err)
	}
	ok, err = CheckAlert("invalid_operator", "1", "<")
	if ok || err == nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">", "invalid", "1")
	if ok || err == nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">", "1", "invalid")
	if ok || err == nil {
		t.Fatal("fail")
	}

	ok, err = CheckAlert(">", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("==", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<=", "1", "0")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("!=", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}

	ok, err = CheckAlert(">", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert(">=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<", "1", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("<=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("==", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlert("!=", "1", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	// check alert string only

	ok, err = CheckAlertOther(">", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther(">=", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("==", "1", "2")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("<=", "1", "0")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("<", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("!=", "1", "1")
	if ok || err != nil {
		t.Fatal("fail")
	}

	ok, err = CheckAlertOther(">", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther(">=", "2", "1")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther(">=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("<", "1", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("<=", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("==", "2", "2")
	if !ok || err != nil {
		t.Fatal("fail")
	}
	ok, err = CheckAlertOther("!=", "1sdasd", "")
	if !ok || err != nil {
		t.Fatal("fail")
	}
}
