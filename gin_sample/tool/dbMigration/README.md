# DB migration

## command

new operation(創建後到 migrations 編輯)

```
# xxx... 可以任意分段, ex:  yarn migrate create init album table
yarn migrate create xxx...
```

migrate db(to new)

```
export DATABASE_URL=<Database Url>
yarn migrate up
```

sample

see `yarn start`

## doc

see `https://salsita.github.io/node-pg-migrate/#/`
