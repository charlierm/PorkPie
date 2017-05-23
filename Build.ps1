go build
New-Item -ItemType Directory -Force -Path .\bin
Compress-Archive -Path .\codedeploy\appspec.yml, .\codedeploy, .\SaturatedFat.exe -DestinationPath .\bin\output.zip -Update