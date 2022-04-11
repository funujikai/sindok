set GOOS=linux
@REM bee pack -ba "-tags prod" -exr="^(?:images|logs|temp|swagger)$"
bee pack -ba "-tags prod"
D:\Project\PuTTY\pscp -pw "!tdJ4y4&j4y42015" Aplikasi.tar.gz itd@10.61.3.87:/home/itd/sindok_dev/api
ssh itd@10.61.3.87