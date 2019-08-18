download_service
===
![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/timoreymann/download_service)

# Beschreibung
Dieser Service dient dem Herunterladen und Bundeln von Remote-Dateien mittels eines kleinen Webservers und Go.

Das Frontend besteht hierbei aus einer einzigen HTML-Datei mit etwas jQuery und AJAX.

# Whats inside?
Ein simpler Service der 1...n Anzahl an Downloads entgegen nimmt, diese asynchron herunterlädt und diese anschliessend gepackt als tar.gz wieder ausspuckt.
Sinnvoll um viele Downloads zu bündeln, Sperren im lokalen Netz zu Umgehen oder multiple Downloads auf einem Server zu bewerkstelligen.

# Usage
Der ganze Service wird via Docker gestartet.

``docker run -p 8086:8086 timoreymann/download_service:latest``

8086 ist hierbei der einzige Port, der freigegeben werden muss. Er stellt den HTTP-Serivce bereit
