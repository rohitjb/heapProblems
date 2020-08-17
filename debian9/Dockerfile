FROM docker.pkg.github.com/anzx/fabric-images/debian9:1595549701

RUN apt-get -qqy update && apt-get install -qqy --no-install-recommends \
        curl \
        gcc \
        apt-transport-https \
        lsb-release \
        gnupg \
        openjdk-8-jdk \
        maven \
        unzip \
    && export CLOUD_SDK_REPO="cloud-sdk-$(lsb_release -c -s)" \
    && echo "deb https://packages.cloud.google.com/apt $CLOUD_SDK_REPO main" > /etc/apt/sources.list.d/google-cloud-sdk.list \
    && curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add - \
    && apt-get update \
    && apt-get install -qqy --no-install-recommends google-cloud-sdk google-cloud-sdk-spanner-emulator google-cloud-sdk-pubsub-emulator \
    && apt-get clean \
    && rm -rf /var/cache/apt/archives

RUN gcloud config configurations create emulator \
    && gcloud config set auth/disable_credentials true \
    && gcloud config set project test-project \
    && gcloud config set api_endpoint_overrides/spanner http://localhost:9020/

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]