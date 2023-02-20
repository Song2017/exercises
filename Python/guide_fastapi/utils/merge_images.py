import base64

import fitz

fedex_path = "/Users/song/_Git/LabSourceCode/Python/guide_fastapi/utils/pdf_base64.txt"
image_path = "/Users/song/_Git/LabSourceCode/Python/guide_fastapi/utils/page.png"

with open(fedex_path, "r") as f:
    delivery_png = base64.b64decode(f.read().encode())
with open(image_path, "rb") as f:
    customs_png = f.read()


# # with open("test.pdf", "w") as f:
# #     f.write(delivery_png.decode())
# # doc = fitz.open(stream=delivery_png, filetype='pdf')
# doc = fitz.Document("pdf", delivery_png)


def append_img_file_to_pdf():
    # Stitch invoice pic to waybill and produce single PDF
    img_doc = fitz.open(image_path)
    img_bytes = img_doc.convert_to_pdf(rotate=90)
    img_pdf = fitz.open("pdf", img_bytes)
    pdf = fitz.Document(
        "pdf",
        base64.b64decode(open(fedex_path, "r").read().encode())
    )
    pdf.insert_pdf(img_pdf)
    _pdf_bytes = pdf.write(deflate=True)
    pdf_base64 = base64.b64encode(_pdf_bytes).decode()
    with open("test.pdf", "wb") as f:
        f.write(_pdf_bytes)
    return pdf_base64


def merge_files_to_pdf():
    img_doc = fitz.open(image_path)
    img_bytes = img_doc.convert_to_pdf(rotate=90)
    # img_pdf = fitz.open("pdf", img_bytes)
    pdf = fitz.Document(
        "pdf",
        base64.b64decode(open(fedex_path, "r").read().encode())
    )
    width, height = fitz.paper_size("a5")
    r = fitz.Rect(0, 0, width, height)
    page = pdf.new_page(-1, width=width, height=height)
    page.insert_image(r, stream=open(image_path, "rb").read())
    _pdf_bytes = pdf.write(deflate=True, garbage=3)
    pdf_base64 = base64.b64encode(_pdf_bytes).decode()
    with open("test2.pdf", "wb") as f:
        f.write(_pdf_bytes)
    return pdf_base64


def append_pdf():
    # with open("test.pdf", "w") as f:
    #     f.write(delivery_png.decode())
    # doc = fitz.open(stream=delivery_png, filetype='pdf')
    doc = fitz.Document("pdf", delivery_png)
    img_doc = fitz.open(image_path)
    img_bytes = img_doc.convert_to_pdf()  # make a 1-page PDF of it
    imgpdf = fitz.open("pdf", img_bytes)
    doc.insert_pdf(imgpdf)
    return doc.write()


def append_page():
    # with open("test.pdf", "w") as f:
    #     f.write(delivery_png.decode())
    # doc = fitz.open(stream=delivery_png, filetype='pdf')
    doc = fitz.Document("pdf", delivery_png)
    img = open(image_path, "rb").read()
    width, height = fitz.paper_size("a4")
    r = fitz.Rect(0, 0, width, height)
    page = doc.new_page(-1, width=width, height=height)
    page.insert_image(r, stream=img)
    return doc.write(encryption=0)


# pdf_bytes = append_pdf(doc)
# import pdb
#
# pdb.set_trace()
# with open("test.pdf", "wb") as f:
#     f.write(pdf_bytes)
# with open("test.txt", "w") as f:
#     f.write(base64.b64encode(pdf_bytes).decode())

merge_files_to_pdf()
# print(merge_files_to_pdf())
