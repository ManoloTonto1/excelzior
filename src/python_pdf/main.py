from win32com import client
import os
from http.server import BaseHTTPRequestHandler, HTTPServer


class handler(BaseHTTPRequestHandler):
    
    def do_POST(self):
        fileBytes = self.rfile.read()

        dir = os.getcwd()

        with open(dir+"/file.xlsx", "wb") as bin_file:
            bin_file.write(fileBytes)


        print("done")
        # Open Microsoft Excel
        excel = client.Dispatch("Excel.Application")
        # Read Excel File
        sheets = excel.Workbooks.Open(dir+'/file.xlsx')
        work_sheets = sheets.Worksheets[0]

        # Convert into PDF File
        try:
            work_sheets.ExportAsFixedFormat(0, dir+'/PDFfile.pdf', 0, 0, 0, 1, 1)
            self.send_response(200)

            self.send_header('Content-type', 'application/pdf')

            self.end_headers()

            with open(dir+"PDFfile.pdf") as file:
                self.wfile.write(file)
        except:
            self.send_response(500)
        finally:
            sheets.Close(True)
            os.remove(dir+"/file.xlsx")
with HTTPServer(('', 8000), handler) as server:
    server.serve_forever()




