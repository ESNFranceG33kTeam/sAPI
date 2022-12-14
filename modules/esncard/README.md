---
title: Module ESNcard
---

# ESNcards endpoint

**Object description**

- Id            int         `json:"id"`             - Id of the esncard
- Id_adherent   int         `json:"id_adherent"`    - Id of the adherent
- Esncard       string      `json:"esncard"`        - Code of the esncard
- CreatedAt     time.Time   `json:"created_at"`     - Created date of the esncard

The parameters `id` and `created_at`are calculated automatically.

## Requirements

This module is dependant of the `adherent` module and need it to works.

## Get

- Get full list of esncards :

```bash
curl -X GET "https://${MYSERVER}/${SECURE}/esncards" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`[json]`) list of esncards objects

- Get only the esncard of a specific adherent :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/esncards/id_adherent/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

- Get only the esncard with the specific esncard code :

```bash
curl -X  GET "https://${MYSERVER}/${SECURE}/esncards/esncard/my-esncard-code" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : (`json`) esncard object

## Post

- To create a new esncard :

```bash
curl -X POST "https://${MYSERVER}/${SECURE}/esncards" \
    -H "accept: application/json" \
    -H "X-Session-Token: ${MYTOKEN}" \
    -H "Content-Type: application/json" \
    -d '{
        "id_adherent": 2,
        "esncard": "aVeryTooLongCode"
    }'
```

Output : (`json`) esncard object

## Delete

- To delete an esncard :

```bash
curl -X DELETE "https://${MYSERVER}/${SECURE}/esncards/3" \
    -H "accept: application/json" -H "X-Session-Token: ${MYTOKEN}"
```

Output : no output

## Data Dependances

In case of suppressing an `adherent` on who has a link an `esncard`, **this `esncard` will be delete as well**.
