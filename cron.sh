#!/bin/bash

syn_api(){
  YEAR=$(gdate -d "+543 years" +%Y)

  SYNAPI=$(curl --location --request POST 'http://localhost:3000/v1/kukr/publishes/sync_datasource' --header 'Content-Type: application/json' --data-raw '{"pub_year": '$YEAR'}')
  echo $SYNAPI
}

syn_api
