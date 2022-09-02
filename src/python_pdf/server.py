from http.server import BaseHTTPRequestHandler, HTTPServer
import os
from subprocess import Popen

serverPort = 8000


class MyServer(BaseHTTPRequestHandler):
    def do_POST(self):
        self.send_response(200)
        fileBytes = self.rfile.read()

        dir = os.getcwd()

        with open(dir+"/file.xlsx", "wb") as bin_file:
            bin_file.write(fileBytes)

        print("done")

        # Convert into PDF File
        try:
            command = 'libreoffice --headless --convert-to pdf ./file.xlsx'
            sts = Popen(command, shell=True).communicate()
            print(sts)
            Popen.kill(sts)
            self.send_response(200)

            self.send_header('Content-type', 'application/pdf')

            self.end_headers()

            with open(dir+"file.pdf") as file:
                self.wfile.write(file)
        except:
            self.send_response(500)
        finally:
            os.remove(dir+"/file.xlsx")


if __name__ == "__main__":
    webServer = HTTPServer(('', serverPort), MyServer)
    print("Server started port 8000")

    try:
        webServer.serve_forever()
    except KeyboardInterrupt:
        pass

    webServer.server_close()
    print("Server stopped.")
