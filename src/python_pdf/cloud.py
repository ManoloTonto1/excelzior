def hello_world(request):
    """Responds to any HTTP request.
    Args:
        request (flask.Request): HTTP request object.
    Returns:
        The response text or any set of values that can be turned into a
        Response object using
        #flask.Flask.make_response>`.
        `make_response <http://flask.pocoo.org/docs/1.0/api/
    """
    fileBytes = request.get_data()

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
            work_sheets.ExportAsFixedFormat(
                0, dir+'/PDFfile.pdf', 0, 0, 0, 1, 1)
            with open(dir+"PDFfile.pdf") as file:
               response = make_response(file)
               response.headers['Content-type'] = 'Content-type'
               response.mimetype = 'application/pdf'
               response.status_code = 200
               return response
        except:
            response = make_response(
                "Error in the creation of the pdf file", 500)
        finally:
            sheets.Close(True)
            os.remove(dir+"/file.xlsx")
