-- +goose Up
CREATETABLEusers(
  idUUIDPRIMARYKEY,
  created_atTIMESTAMPNOTNULL,
  updated_atTIMESTAMPNOTNULL,
  nameTEXTNOTNULL
);
-- +goose Down
DROPTABLEusers;
