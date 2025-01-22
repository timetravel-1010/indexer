## Run the project

Start creating a data folder to store the ZincSearch data

```bash
mkdir data
```

Then, grant the necessary permissions

```bash
chmod -R a+rwx data
```

Finally you can run docker compose

```bash
docker compose up --build
```

Now you have the following services running on your machine:

- Go API running on port 8080
- Web Client running on port 3000
- ZincSearch running on port 4080
