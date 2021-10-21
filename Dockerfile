FROM suborbital/atmo:v0.3.2

COPY ./runnables.wasm.zip .

ENTRYPOINT [ "atmo" ]
