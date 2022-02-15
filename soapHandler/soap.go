package soapHandler

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Request struct {
	//Values are set in below fields as per the request
	CodigoProducto string
	CodigoSociedad string
	CodigoSede     string
}

type Response struct {
	XMLName  xml.Name          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope" json:"-"`
	SoapBody *SOAPBodyResponse `xml:"Body" json:"Body"`
}

type SOAPBodyResponse struct {
	Resp         *GetStockResponseBody `xml:"ZSDRFC_SKN_GET_STOCKResponse" json:"Response"`
	FaultDetails *Fault                `xml:"Fault" json:"fault,omitempty"`
}

type Fault struct {
	Faultcode   string `xml:"faultcode"`
	Faultstring string `xml:"faultstring"`
}

type GetStockResponseBody struct {
	Result *Return `xml:"ET_RETURN"`
	Stock  *Stock  `xml:"ET_STOCK"`
}

type Return struct {
	Item *ReturnItem `xml:"item"`
}

type ReturnItem struct {
	Type          string `xml:"TYPE" json:"type"`
	Code          string `xml:"CODE" json:"code"`
	ResultMessage string `xml:"MESSAGE"`
	Log_No        string `xml:"LOG_NO" json:"log_no,omitempty"`
	LOG_MSG_NO    string `xml:"LOG_MSG_NO"`
	MESSAGE_V1    string `xml:"MESSAGE_V1" json:"message1,omitempty"`
	MESSAGE_V2    string `xml:"MESSAGE_V2" json:"message2,omitempty"`
	MESSAGE_V3    string `xml:"MESSAGE_V3" json:"message3,omitempty"`
	MESSAGE_V4    string `xml:"MESSAGE_V4" json:"message4,omitempty"`
}

type Stock struct {
	Product []*StockItem `xml:"item" json:"product`
}

type StockItem struct {
	ProductId               string `xml:"ZSD_CMATER" json:"id"`
	Description             string `xml:"ZSD_DCORTA" json:"description"`
	CommercialQuantity      string `xml:"ZSD_QSRUCO" json:"commercialQty"`
	CommercialUnitOfMeasure string `xml:"ZSD_CUMUCO" json:"commerciaUnitOfMeasure"`
	BaseQuantity            string `xml:"ZSD_QSRBAS" json:"baseQty"`
	BaseUnitOfMeasure       string `xml:"ZSD_CUMBAS" json:"baseUnitOfMeasure"`
}

func generateSOAPRequest(req *Request) (*http.Request, error) {

	var requestTemplate = getRequestTemplate(req)

	// Using the var getTemplate to construct request
	template, err := template.New("InputRequest").Parse(requestTemplate)
	if err != nil {
		fmt.Printf("Error while marshling object. %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Printf("template.Execute error. %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Printf("encoder.Encode error. %s ", err.Error())
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, "https://servicioswebdex.alicorp.com.pe/nd1/sap/bc/srt/rfc/sap/zsdrfc_skn_get_stock/300/zsdrfc_skn_get_stock/zsdrfc_skn_get_stock", bytes.NewBuffer(doc.Bytes()))
	r.SetBasicAuth("nrodriguezv", "mikaela2013")
	r.Header.Add("Content-Type", "text/xml;charset=UTF-8")

	if err != nil {
		fmt.Printf("Error making a request. %s ", err.Error())
		return nil, err
	}

	return r, nil
}

func CallSOAPClientSteps(req *Request) (*Response, error) {

	httpReq, err := generateSOAPRequest(req)
	if err != nil {
		fmt.Println("Some problem occurred in request generation")
	}

	response, err := soapCall(httpReq)
	if err != nil {
		fmt.Println("Problem occurred in making a SOAP call")
	}

	return response, err

}

func soapCall(req *http.Request) (*Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r := &Response{}
	err = xml.Unmarshal(body, &r)

	if err != nil {
		return nil, err
	}

	//if r.SoapBody.Resp.Status != "200" {
	//	return nil, err
	//}

	return r, nil
}
