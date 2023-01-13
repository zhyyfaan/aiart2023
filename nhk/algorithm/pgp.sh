#! /bin/bash

input_dir=$1;
output_dir=$2;

echo ${input_dir} ${output_dir}

cp ${input_dir} ${output_dir};

echo "pgp算法运行结束" >> log.log;
exit 0;