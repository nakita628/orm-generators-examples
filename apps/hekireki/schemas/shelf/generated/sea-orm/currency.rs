use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, EnumIter, DeriveActiveEnum, Serialize, Deserialize)]
#[sea_orm(rs_type = "String", db_type = "Enum", enum_name = "Currency")]
pub enum Currency {
    #[sea_orm(string_value = "AED")]
    Aed,
    #[sea_orm(string_value = "AFN")]
    Afn,
    #[sea_orm(string_value = "ALL")]
    All,
    #[sea_orm(string_value = "AMD")]
    Amd,
    #[sea_orm(string_value = "ANG")]
    Ang,
    #[sea_orm(string_value = "AOA")]
    Aoa,
    #[sea_orm(string_value = "ARS")]
    Ars,
    #[sea_orm(string_value = "AUD")]
    Aud,
    #[sea_orm(string_value = "AWG")]
    Awg,
    #[sea_orm(string_value = "AZN")]
    Azn,
    #[sea_orm(string_value = "BAM")]
    Bam,
    #[sea_orm(string_value = "BBD")]
    Bbd,
    #[sea_orm(string_value = "BDT")]
    Bdt,
    #[sea_orm(string_value = "BGN")]
    Bgn,
    #[sea_orm(string_value = "BHD")]
    Bhd,
    #[sea_orm(string_value = "BIF")]
    Bif,
    #[sea_orm(string_value = "BMD")]
    Bmd,
    #[sea_orm(string_value = "BND")]
    Bnd,
    #[sea_orm(string_value = "BOB")]
    Bob,
    #[sea_orm(string_value = "BRL")]
    Brl,
    #[sea_orm(string_value = "BSD")]
    Bsd,
    #[sea_orm(string_value = "BTN")]
    Btn,
    #[sea_orm(string_value = "BWP")]
    Bwp,
    #[sea_orm(string_value = "BYN")]
    Byn,
    #[sea_orm(string_value = "BZD")]
    Bzd,
    #[sea_orm(string_value = "CAD")]
    Cad,
    #[sea_orm(string_value = "CDF")]
    Cdf,
    #[sea_orm(string_value = "CHF")]
    Chf,
    #[sea_orm(string_value = "CLP")]
    Clp,
    #[sea_orm(string_value = "CNY")]
    Cny,
    #[sea_orm(string_value = "COP")]
    Cop,
    #[sea_orm(string_value = "CRC")]
    Crc,
    #[sea_orm(string_value = "CUP")]
    Cup,
    #[sea_orm(string_value = "CVE")]
    Cve,
    #[sea_orm(string_value = "CZK")]
    Czk,
    #[sea_orm(string_value = "DJF")]
    Djf,
    #[sea_orm(string_value = "DKK")]
    Dkk,
    #[sea_orm(string_value = "DOP")]
    Dop,
    #[sea_orm(string_value = "DZD")]
    Dzd,
    #[sea_orm(string_value = "EGP")]
    Egp,
    #[sea_orm(string_value = "ERN")]
    Ern,
    #[sea_orm(string_value = "ETB")]
    Etb,
    #[sea_orm(string_value = "EUR")]
    Eur,
    #[sea_orm(string_value = "FJD")]
    Fjd,
    #[sea_orm(string_value = "FKP")]
    Fkp,
    #[sea_orm(string_value = "GBP")]
    Gbp,
    #[sea_orm(string_value = "GEL")]
    Gel,
    #[sea_orm(string_value = "GHS")]
    Ghs,
    #[sea_orm(string_value = "GIP")]
    Gip,
    #[sea_orm(string_value = "GMD")]
    Gmd,
    #[sea_orm(string_value = "GNF")]
    Gnf,
    #[sea_orm(string_value = "GTQ")]
    Gtq,
    #[sea_orm(string_value = "GYD")]
    Gyd,
    #[sea_orm(string_value = "HKD")]
    Hkd,
    #[sea_orm(string_value = "HNL")]
    Hnl,
    #[sea_orm(string_value = "HTG")]
    Htg,
    #[sea_orm(string_value = "HUF")]
    Huf,
    #[sea_orm(string_value = "IDR")]
    Idr,
    #[sea_orm(string_value = "ILS")]
    Ils,
    #[sea_orm(string_value = "INR")]
    Inr,
    #[sea_orm(string_value = "IQD")]
    Iqd,
    #[sea_orm(string_value = "IRR")]
    Irr,
    #[sea_orm(string_value = "ISK")]
    Isk,
    #[sea_orm(string_value = "JMD")]
    Jmd,
    #[sea_orm(string_value = "JOD")]
    Jod,
    #[sea_orm(string_value = "JPY")]
    Jpy,
    #[sea_orm(string_value = "KES")]
    Kes,
    #[sea_orm(string_value = "KGS")]
    Kgs,
    #[sea_orm(string_value = "KHR")]
    Khr,
    #[sea_orm(string_value = "KMF")]
    Kmf,
    #[sea_orm(string_value = "KPW")]
    Kpw,
    #[sea_orm(string_value = "KRW")]
    Krw,
    #[sea_orm(string_value = "KWD")]
    Kwd,
    #[sea_orm(string_value = "KYD")]
    Kyd,
    #[sea_orm(string_value = "KZT")]
    Kzt,
    #[sea_orm(string_value = "LAK")]
    Lak,
    #[sea_orm(string_value = "LBP")]
    Lbp,
    #[sea_orm(string_value = "LKR")]
    Lkr,
    #[sea_orm(string_value = "LRD")]
    Lrd,
    #[sea_orm(string_value = "LSL")]
    Lsl,
    #[sea_orm(string_value = "LYD")]
    Lyd,
    #[sea_orm(string_value = "MAD")]
    Mad,
    #[sea_orm(string_value = "MDL")]
    Mdl,
    #[sea_orm(string_value = "MGA")]
    Mga,
    #[sea_orm(string_value = "MKD")]
    Mkd,
    #[sea_orm(string_value = "MMK")]
    Mmk,
    #[sea_orm(string_value = "MNT")]
    Mnt,
    #[sea_orm(string_value = "MOP")]
    Mop,
    #[sea_orm(string_value = "MRU")]
    Mru,
    #[sea_orm(string_value = "MUR")]
    Mur,
    #[sea_orm(string_value = "MVR")]
    Mvr,
    #[sea_orm(string_value = "MWK")]
    Mwk,
    #[sea_orm(string_value = "MXN")]
    Mxn,
    #[sea_orm(string_value = "MYR")]
    Myr,
    #[sea_orm(string_value = "MZN")]
    Mzn,
    #[sea_orm(string_value = "NAD")]
    Nad,
    #[sea_orm(string_value = "NGN")]
    Ngn,
    #[sea_orm(string_value = "NIO")]
    Nio,
    #[sea_orm(string_value = "NOK")]
    Nok,
    #[sea_orm(string_value = "NPR")]
    Npr,
    #[sea_orm(string_value = "NZD")]
    Nzd,
    #[sea_orm(string_value = "OMR")]
    Omr,
    #[sea_orm(string_value = "PAB")]
    Pab,
    #[sea_orm(string_value = "PEN")]
    Pen,
    #[sea_orm(string_value = "PGK")]
    Pgk,
    #[sea_orm(string_value = "PHP")]
    Php,
    #[sea_orm(string_value = "PKR")]
    Pkr,
    #[sea_orm(string_value = "PLN")]
    Pln,
    #[sea_orm(string_value = "PYG")]
    Pyg,
    #[sea_orm(string_value = "QAR")]
    Qar,
    #[sea_orm(string_value = "RON")]
    Ron,
    #[sea_orm(string_value = "RSD")]
    Rsd,
    #[sea_orm(string_value = "RUB")]
    Rub,
    #[sea_orm(string_value = "RWF")]
    Rwf,
    #[sea_orm(string_value = "SAR")]
    Sar,
    #[sea_orm(string_value = "SBD")]
    Sbd,
    #[sea_orm(string_value = "SCR")]
    Scr,
    #[sea_orm(string_value = "SDG")]
    Sdg,
    #[sea_orm(string_value = "SEK")]
    Sek,
    #[sea_orm(string_value = "SGD")]
    Sgd,
    #[sea_orm(string_value = "SHP")]
    Shp,
    #[sea_orm(string_value = "SLE")]
    Sle,
    #[sea_orm(string_value = "SOS")]
    Sos,
    #[sea_orm(string_value = "SRD")]
    Srd,
    #[sea_orm(string_value = "SSP")]
    Ssp,
    #[sea_orm(string_value = "STN")]
    Stn,
    #[sea_orm(string_value = "SVC")]
    Svc,
    #[sea_orm(string_value = "SYP")]
    Syp,
    #[sea_orm(string_value = "SZL")]
    Szl,
    #[sea_orm(string_value = "THB")]
    Thb,
    #[sea_orm(string_value = "TJS")]
    Tjs,
    #[sea_orm(string_value = "TMT")]
    Tmt,
    #[sea_orm(string_value = "TND")]
    Tnd,
    #[sea_orm(string_value = "TOP")]
    Top,
    #[sea_orm(string_value = "TRY")]
    Try,
    #[sea_orm(string_value = "TTD")]
    Ttd,
    #[sea_orm(string_value = "TWD")]
    Twd,
    #[sea_orm(string_value = "TZS")]
    Tzs,
    #[sea_orm(string_value = "UAH")]
    Uah,
    #[sea_orm(string_value = "UGX")]
    Ugx,
    #[sea_orm(string_value = "USD")]
    Usd,
    #[sea_orm(string_value = "UYU")]
    Uyu,
    #[sea_orm(string_value = "UZS")]
    Uzs,
    #[sea_orm(string_value = "VES")]
    Ves,
    #[sea_orm(string_value = "VND")]
    Vnd,
    #[sea_orm(string_value = "VUV")]
    Vuv,
    #[sea_orm(string_value = "WST")]
    Wst,
    #[sea_orm(string_value = "XAF")]
    Xaf,
    #[sea_orm(string_value = "XCD")]
    Xcd,
    #[sea_orm(string_value = "XOF")]
    Xof,
    #[sea_orm(string_value = "XPF")]
    Xpf,
    #[sea_orm(string_value = "YER")]
    Yer,
    #[sea_orm(string_value = "ZAR")]
    Zar,
    #[sea_orm(string_value = "ZMW")]
    Zmw,
    #[sea_orm(string_value = "ZWL")]
    Zwl,
}
