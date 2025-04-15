This project processes the NSEScripMaster.txt file provided by the [ICICI Direct Breeze API](https://api.icicidirect.com/breezeapi/documents/index.html#instruments)

The raw .txt file was converted into a structured CSV format.

Each row from the CSV was stored as an individual document in MongoDB for faster search and retrieval.

This setup enables efficient querying of instrument metadata by symbol, token, or other identifiers.
