package soapHandler

func getRequestTemplate(req *Request) string {
   request := req

   if len(request.CodigoProducto) > 0 {
      return getTemplate
   } else {
      return getTemplateAllProducts
   }
}

var getTemplate = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:sap-com:document:sap:rfc:functions">
   <soapenv:Header/>
   <soapenv:Body>
      <urn:ZSDRFC_SKN_GET_STOCK>
         <ET_RETURN>
         </ET_RETURN>
         <ET_STOCK>
         </ET_STOCK>
         <IT_FILTROS>
            <item>
               <FIELDNM>COD_SOCIEDAD</FIELDNM>
               <SIGN>I</SIGN>
               <OPTION>EQ</OPTION>
               <LOW>{{.CodigoSociedad}}</LOW>
               <HIGH></HIGH>
            </item>
            <item>
               <FIELDNM>COD_SEDE</FIELDNM>
               <SIGN>I</SIGN>
               <OPTION>EQ</OPTION>
               <LOW>{{.CodigoSede}}</LOW>
               <HIGH></HIGH>
            </item>
            <item>
               <FIELDNM>COD_PRODUCTO</FIELDNM>
               <SIGN>I</SIGN>
               <OPTION>EQ</OPTION>
               <LOW>{{.CodigoProducto}}</LOW>
               <HIGH></HIGH>
            </item>
         </IT_FILTROS>
      </urn:ZSDRFC_SKN_GET_STOCK>
   </soapenv:Body>
</soapenv:Envelope>
`
var getTemplateAllProducts = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:sap-com:document:sap:rfc:functions">
   <soapenv:Header/>
   <soapenv:Body>
      <urn:ZSDRFC_SKN_GET_STOCK>
         <ET_RETURN>
         </ET_RETURN>
         <ET_STOCK>
         </ET_STOCK>
         <IT_FILTROS>
            <item>
               <FIELDNM>COD_SOCIEDAD</FIELDNM>
               <SIGN>I</SIGN>
               <OPTION>EQ</OPTION>
               <LOW>{{.CodigoSociedad}}</LOW>
               <HIGH></HIGH>
            </item>
            <item>
               <FIELDNM>COD_SEDE</FIELDNM>
               <SIGN>I</SIGN>
               <OPTION>EQ</OPTION>
               <LOW>{{.CodigoSede}}</LOW>
               <HIGH></HIGH>
            </item>
         </IT_FILTROS>
      </urn:ZSDRFC_SKN_GET_STOCK>
   </soapenv:Body>
</soapenv:Envelope>
`


/*
<soap-env:Envelope xmlns:soap-env="http://schemas.xmlsoap.org/soap/envelope/"
   <soap-env:Header/>
   <soap-env:Body>
      <n0:ZSDRFC_SKN_GET_STOCKResponse xmlns:n0="urn:sap-com:document:sap:rfc:functions"
         <ET_RETURN>
            <item>
               <TYPE>S</TYPE
               <CODE>01</CODE
               <MESSAGE>Datos obtenidos exitosamente.</MESSAGE
               <LOG_NO/
               <LOG_MSG_NO>000000</LOG_MSG_NO
               <MESSAGE_V1/
               <MESSAGE_V2/>
               <MESSAGE_V3/>
               <MESSAGE_V4/>
            </item>
         </ET_RETURN>
         <ET_STOCK>
            <item>
               <ZSD_CSOCIE>801</ZSD_CSOCIE
               <ZSD_CSEDE>0048</ZSD_CSEDE>
               <ZSD_CMATER>000000000000018110</ZSD_CMATER
               <ZSD_DCORTA>HARINA BLANCA FLOR PREP.1K 12BOL</ZSD_DCORTA
               <ZSD_QSTUCO>574.0</ZSD_QSTUCO
               <ZSD_QSTBAS>3.0</ZSD_QSTBAS
               <ZSD_QPVUCO>0.0</ZSD_QPVUCO
               <ZSD_QPVBAS>0.0</ZSD_QPVBAS
               <ZSD_QSRUCO>574.0</ZSD_QSRUCO
               <ZSD_QSRBAS>3.0</ZSD_QSRBAS
               <ZSD_CUMUCO>PQT</ZSD_CUMUCO>
               <ZSD_CUMBAS>BOL</ZSD_CUMBAS
               <ZSD_NFUVTA>12</ZSD_NFUVTA
            </item>
         </ET_STOCK>
         <IT_FILTROS
            <item
               <FIELDNM>COD_SOCIEDAD</FIELDNM
               <SIGN>I</SIGN
               <OPTION>EQ</OPTION
               <LOW>801</LOW
               <HIGH/
            </item
            <item>
               <FIELDNM>COD_SEDE</FIELDNM>
               <SIGN>I</SIGN
               <OPTION>EQ</OPTION
               <LOW>0048</LOW>
               <HIGH/
            </item
            <item>
               <FIELDNM>COD_PRODUCTO</FIELDNM
               <SIGN>I</SIGN
               <OPTION>EQ</OPTION
               <LOW>18110</LOW
               <HIGH/>
            </item>
         </IT_FILTROS>
      </n0:ZSDRFC_SKN_GET_STOCKResponse>
   </soap-env:Body>
</soap-env:Envelope>
*/
