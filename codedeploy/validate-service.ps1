Invoke-WebRequest -UseBasicParsing -Uri healthcheck.local | ForEach-Object {if ($_.StatusCode -ne 200){throw "not up"} }