package pkg

import (
	"fmt"
	"testing"
)

func Test_MustGet(t *testing.T) {
	text := `{
  "SYS_HEAD": {
    "TRAN_TIMESTAMP": "1626409557477",
    "SOURCE_BRANCH_NO": "999",
    "RET_STATUS": "S",
    "SEQ_NO": "SEQNO-SMOKE-573e751ed625",
    "SOURCE_TYPE": "MB",
    "REFERENCE": "ENS2021060400006001",
    "MESSAGE_CODE": "0007",
    "SERVICE_CODE": "MbsdCore",
    "COMPANY": "",
    "RET": [
      {
        "RET_MSG": "000000 SUCCESS",
        "RET_CODE": "000000"
      }
    ],
    "SCENE_ID": "01",
    "DEST_BRANCH_NO": "999",
    "TRAN_DATE": "20210604",
    "AUTH_USER_ID": "",
    "MESSAGE_TYPE": "1400",
    "RUN_DATE": "20210604"
  },
  "BODY": {
    "CONTACT_TEL": "62TEL-SMOKE-13153078336",
    "SETTLE_ARRAY": [],
    "CLIENT_STATUS": "A",
    "ANNUAL_DATE": "20210604",
    "NATURE_CLASS": "V",
    "ALT_ACCT_NAME": "smoke test",
    "FACE_FLAG": "1",
    "GL_TYPE": "R",
    "ACCT_DESC": "All-in-one (AIO account)",
    "INT_IND": "N",
    "WITHDRAWAL_TYPE": "P",
    "SOURCE_MODULE": "RB",
    "AUTO_DEP": "N",
    "APPLY_BRANCH": "999",
    "CLIENT_NO": "100000202",
    "LEAD_ACCT_FLAG": "Y",
    "ACCT_STATUS": "A",
    "IS_INDIVIDUAL": "Y",
    "FIN_BAL": 996999,
    "EFFECT_DATE": "20210604",
    "ISS_COUNTRY": "ID",
    "JUDICIAL_FREEZE_FLAG": "N",
    "ANNUAL_STATUS": "Y",
    "DOCUMENT_ID": "KTP-SMOKE-313bafd4aaab",
    "ACCT_EXEC": "",
    "ACCT_REAL_FLAG": "N",
    "INTERNAL_KEY": "205000",
    "CONTACT_ARRAY": [
      {
        "LINKMAN_TYPE": "05",
        "DOCUMENT_ID": "",
        "PHONE_NO1": "",
        "DOCUMENT_TYPE": "",
        "PHONE_NO2": "",
        "LINKMAN_NAME": ""
      }
    ],
    "ACCT_NAME": "smoke test",
    "LAST_CHANGE_USER_ID": "MBU001",
    "DOCUMENT_TYPE": "101",
    "CLIENT_TYPE": "01",
    "ACCT_TYPE": "A",
    "ACCT_SEQ_NO": "0",
    "CLIENT_INDICATOR": "N",
    "IS_STOP_PAY": "N",
    "CCY": "IDR",
    "SOURCE_TYPE": "MB",
    "FTA_FLAG": "N",
    "BAL_TYPE": "CA",
    "CLIENT_NAME": "DEBBY1 ANGGRAINI2",
    "ALL_DRA_IND": "Y",
    "BASE_ACCT_NO": "901508344410",
    "USER_ID": "MBU001",
    "LAST_CHANGE_DATE": "20210604",
    "HOME_BRANCH": "999",
    "DAY_NUM": "0",
    "CA_TT_FLAG": "0",
    "ACCOUNTING_STATUS": "ZHC",
    "BRANCH": "999",
    "PAY_PWD_IND": "N",
    "PROD_TYPE": "901",
    "PROFIT_CENTRE": "99",
    "ACCT_NATURE": "20",
    "ACCT_OPEN_DATE": "20210604",
    "ALL_DEP_IND": "Y"
  },
  "APP_HEAD": {
    "PAGE_END": "0",
    "CURRENT_NUM": "0",
    "TOTAL_FLAG": "E",
    "PGUP_OR_PGDN": "0",
    "PAGE_START": "0",
    "TOTAL_NUM": "-1"
  }
}`

	obj := MustGet(text, []interface{}{"SYS_HEAD", "RET", 0, "RET_MSG"})
	fmt.Println(obj)

	obj = MustGet(text, []interface{}{"BODY", "CONTACT_TEL"})
	fmt.Println(obj)
}
