import os
from http.server import BaseHTTPRequestHandler, HTTPServer
from subprocess import Popen
from sys import stderr
import threading




class handler(BaseHTTPRequestHandler):
    def runCommand():


    def do_POST(self):
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


server = HTTPServer(('', 8080), handler)
server.serve_forever()


