FROM gcr.io/cloud-marketplace-containers/google/debian9:latest@sha256:fc35333ac79fa4ea521cb0656642b1bf5ba07742f5b04c2ae79d3a934e85f21f

# Add ANZ CAs
COPY ./certs/*.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

ENV ANZ="true"
