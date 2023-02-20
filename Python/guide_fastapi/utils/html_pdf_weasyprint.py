from weasyprint import HTML

html = """
<html>
<meta charset="UTF-8">
<title>Invocie</title>
<style type="text/css">
    html,
    body {
        width: 595pt;
        height: 842pt;
    }

    #tableBolder,
    #tableBolder td {
        border: 1px solid black;
        border-collapse: collapse;
    }

    #tableBolder p {
        white-space: nowrap;
        font-size: x-small;
        margin-right: 5px;
    }

    .totalSpan {
        display: inline-block;
        width: 180px;
    }
</style>

<body style="background-color: white; font-family: ui-monospace">
    <table style="text-align: left; font-weight: bold;">
        <tr style="height: 50px;"></tr>
        <tr>
            <td colspan="2">
                <span style="font-size: x-large;">Commercial Invoice</span>
            </td>
        </tr>
        <tr>
            <td colspan="2">Invoice No.:<span style="margin-left: 20px;">invoice_no</span>
            </td>
        </tr>
        <tr>
            <td style="width: 40%;">Shipper</td>
            <td>Buyer</td>
        </tr>
        <tr>
            <td>shipper_company</td>
            <td>buyer_company</td>
        </tr>
        <tr>
            <td>shipper_address</td>
            <td>buyer_address</td>
        </tr>
        <tr>
            <td>shipper_post_code</td>
            <td>buyer_post_code</td>
        </tr>
        <tr>
            <td>shipper_phone</td>
            <td>buyer_phone</td>
        </tr>
        <tr>
            <td>shipper_coutry</td>
            <td>buyer_coutry</td>
        </tr>
        <tr>
            <td colspan="2"><span class="totalSpan">Total Value(RMB):</span><span style="color: red;">total_value</span>
            </td>
        </tr>
        <tr>
            <td colspan="2"><span class="totalSpan">Total Boxes/Pallets:</span><span>total_boxes</span></td>
        </tr>
        <tr>
            <td colspan="2"><span class="totalSpan">Total Weight(KG):</span><span>total_weight</span></td>
        </tr>
        <tr style="height: 5px;"></tr>
        <tr>
            <td colspan="2">
                <table id="tableBolder">
                    <tr>
                        <td><p>Parcel weight</p></td>
                        <td><p>Band name1</p></td>
                        <td><p style="margin-right: 50px;">Product name</p></td>
                        <td><p>Quantity 1</p></td>
                        <td><p>Unit Price 1</p></td>
                        <td><p>HS Code</p></td>
                        <td><p>Barcod</p></td>
                        <td><p>Tax</p></td>
                        <td><p>Shipping Cost</p></td>
                        <td><p>Total Price</p></td>
                        <td><p>Actual Paid</p></td>
                    </tr>
                    <tr>
                        <td><p>parcel_weight</p></td>
                        <td><p>band_name</p></td>
                        <td><p>product_name</p></td>
                        <td><p>quantity</p></td>
                        <td><p>unit_price</p></td>
                        <td><p>hs_code</p></td>
                        <td><p>barcod</p></td>
                        <td><p>tax</p></td>
                        <td><p>shipping_cost</p></td>
                        <td><p>total_price</p></td>
                        <td><p>actual_paid</p></td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>

</html>
"""
HTML(string=html).write_pdf('path_of_pdf.pdf')
