# devops-netology

### My first repo

1. В корневом каталоге ни один файл не игнорируется
2. В каталоге [terraform](./terraform) игнорируются:
   + директории .terraform, включая вложенные.
   + файлы с расширением .tfstate и любыми символами, предварёнными точкой, после него
   + файлы crash.log 
   + файлы с расширением .tfvars
   + файлы override.tf, override.tf.json и другие, заканчивающиеся на _override.tf и _override.tf.json
   + конфигурационные файлы .terraformrc и terraform.rc

