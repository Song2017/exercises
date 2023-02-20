1. Gen Code of Model 
- https://github.com/smallnest/gen
```
go install github.com/smallnest/gen@latest     

~/go/bin/gen --sqltype=postgres \
   	--connstr "postgresql://nomad_logistics_dev:mIPqf6kC0eqinBOFevwvgArLc0I487@pgm-6nn7xj3p540344s2167570.pg.rds.aliyuncs.com/nomad_logistics_dev?sslmode=disable" \
   	--database main  \
    --table fc_configuration_template \
   	--json \
   	--gorm \
   	--guregu \
   	--rest \
   	--out ./example \
   	--module example.com/rest/example \
   	--mod \
   	--server \
   	--makefile \
   	--json-fmt=snake \
   	--generate-dao \
   	--generate-proj \
   	--overwrite
```
