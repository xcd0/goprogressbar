#!/bin/bash

max=1000
for ((i=0; i < max; i+=2))
do
	./goprogress $i
done

