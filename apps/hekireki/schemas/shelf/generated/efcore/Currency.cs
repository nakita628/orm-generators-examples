#nullable enable

using NpgsqlTypes;

namespace Shelf.Models;

public enum Currency
{
    [PgName("AED")]
    Aed,

    [PgName("AFN")]
    Afn,

    [PgName("ALL")]
    All,

    [PgName("AMD")]
    Amd,

    [PgName("ANG")]
    Ang,

    [PgName("AOA")]
    Aoa,

    [PgName("ARS")]
    Ars,

    [PgName("AUD")]
    Aud,

    [PgName("AWG")]
    Awg,

    [PgName("AZN")]
    Azn,

    [PgName("BAM")]
    Bam,

    [PgName("BBD")]
    Bbd,

    [PgName("BDT")]
    Bdt,

    [PgName("BGN")]
    Bgn,

    [PgName("BHD")]
    Bhd,

    [PgName("BIF")]
    Bif,

    [PgName("BMD")]
    Bmd,

    [PgName("BND")]
    Bnd,

    [PgName("BOB")]
    Bob,

    [PgName("BRL")]
    Brl,

    [PgName("BSD")]
    Bsd,

    [PgName("BTN")]
    Btn,

    [PgName("BWP")]
    Bwp,

    [PgName("BYN")]
    Byn,

    [PgName("BZD")]
    Bzd,

    [PgName("CAD")]
    Cad,

    [PgName("CDF")]
    Cdf,

    [PgName("CHF")]
    Chf,

    [PgName("CLP")]
    Clp,

    [PgName("CNY")]
    Cny,

    [PgName("COP")]
    Cop,

    [PgName("CRC")]
    Crc,

    [PgName("CUP")]
    Cup,

    [PgName("CVE")]
    Cve,

    [PgName("CZK")]
    Czk,

    [PgName("DJF")]
    Djf,

    [PgName("DKK")]
    Dkk,

    [PgName("DOP")]
    Dop,

    [PgName("DZD")]
    Dzd,

    [PgName("EGP")]
    Egp,

    [PgName("ERN")]
    Ern,

    [PgName("ETB")]
    Etb,

    [PgName("EUR")]
    Eur,

    [PgName("FJD")]
    Fjd,

    [PgName("FKP")]
    Fkp,

    [PgName("GBP")]
    Gbp,

    [PgName("GEL")]
    Gel,

    [PgName("GHS")]
    Ghs,

    [PgName("GIP")]
    Gip,

    [PgName("GMD")]
    Gmd,

    [PgName("GNF")]
    Gnf,

    [PgName("GTQ")]
    Gtq,

    [PgName("GYD")]
    Gyd,

    [PgName("HKD")]
    Hkd,

    [PgName("HNL")]
    Hnl,

    [PgName("HTG")]
    Htg,

    [PgName("HUF")]
    Huf,

    [PgName("IDR")]
    Idr,

    [PgName("ILS")]
    Ils,

    [PgName("INR")]
    Inr,

    [PgName("IQD")]
    Iqd,

    [PgName("IRR")]
    Irr,

    [PgName("ISK")]
    Isk,

    [PgName("JMD")]
    Jmd,

    [PgName("JOD")]
    Jod,

    [PgName("JPY")]
    Jpy,

    [PgName("KES")]
    Kes,

    [PgName("KGS")]
    Kgs,

    [PgName("KHR")]
    Khr,

    [PgName("KMF")]
    Kmf,

    [PgName("KPW")]
    Kpw,

    [PgName("KRW")]
    Krw,

    [PgName("KWD")]
    Kwd,

    [PgName("KYD")]
    Kyd,

    [PgName("KZT")]
    Kzt,

    [PgName("LAK")]
    Lak,

    [PgName("LBP")]
    Lbp,

    [PgName("LKR")]
    Lkr,

    [PgName("LRD")]
    Lrd,

    [PgName("LSL")]
    Lsl,

    [PgName("LYD")]
    Lyd,

    [PgName("MAD")]
    Mad,

    [PgName("MDL")]
    Mdl,

    [PgName("MGA")]
    Mga,

    [PgName("MKD")]
    Mkd,

    [PgName("MMK")]
    Mmk,

    [PgName("MNT")]
    Mnt,

    [PgName("MOP")]
    Mop,

    [PgName("MRU")]
    Mru,

    [PgName("MUR")]
    Mur,

    [PgName("MVR")]
    Mvr,

    [PgName("MWK")]
    Mwk,

    [PgName("MXN")]
    Mxn,

    [PgName("MYR")]
    Myr,

    [PgName("MZN")]
    Mzn,

    [PgName("NAD")]
    Nad,

    [PgName("NGN")]
    Ngn,

    [PgName("NIO")]
    Nio,

    [PgName("NOK")]
    Nok,

    [PgName("NPR")]
    Npr,

    [PgName("NZD")]
    Nzd,

    [PgName("OMR")]
    Omr,

    [PgName("PAB")]
    Pab,

    [PgName("PEN")]
    Pen,

    [PgName("PGK")]
    Pgk,

    [PgName("PHP")]
    Php,

    [PgName("PKR")]
    Pkr,

    [PgName("PLN")]
    Pln,

    [PgName("PYG")]
    Pyg,

    [PgName("QAR")]
    Qar,

    [PgName("RON")]
    Ron,

    [PgName("RSD")]
    Rsd,

    [PgName("RUB")]
    Rub,

    [PgName("RWF")]
    Rwf,

    [PgName("SAR")]
    Sar,

    [PgName("SBD")]
    Sbd,

    [PgName("SCR")]
    Scr,

    [PgName("SDG")]
    Sdg,

    [PgName("SEK")]
    Sek,

    [PgName("SGD")]
    Sgd,

    [PgName("SHP")]
    Shp,

    [PgName("SLE")]
    Sle,

    [PgName("SOS")]
    Sos,

    [PgName("SRD")]
    Srd,

    [PgName("SSP")]
    Ssp,

    [PgName("STN")]
    Stn,

    [PgName("SVC")]
    Svc,

    [PgName("SYP")]
    Syp,

    [PgName("SZL")]
    Szl,

    [PgName("THB")]
    Thb,

    [PgName("TJS")]
    Tjs,

    [PgName("TMT")]
    Tmt,

    [PgName("TND")]
    Tnd,

    [PgName("TOP")]
    Top,

    [PgName("TRY")]
    Try,

    [PgName("TTD")]
    Ttd,

    [PgName("TWD")]
    Twd,

    [PgName("TZS")]
    Tzs,

    [PgName("UAH")]
    Uah,

    [PgName("UGX")]
    Ugx,

    [PgName("USD")]
    Usd,

    [PgName("UYU")]
    Uyu,

    [PgName("UZS")]
    Uzs,

    [PgName("VES")]
    Ves,

    [PgName("VND")]
    Vnd,

    [PgName("VUV")]
    Vuv,

    [PgName("WST")]
    Wst,

    [PgName("XAF")]
    Xaf,

    [PgName("XCD")]
    Xcd,

    [PgName("XOF")]
    Xof,

    [PgName("XPF")]
    Xpf,

    [PgName("YER")]
    Yer,

    [PgName("ZAR")]
    Zar,

    [PgName("ZMW")]
    Zmw,

    [PgName("ZWL")]
    Zwl,
}
