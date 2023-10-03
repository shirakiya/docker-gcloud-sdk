FROM python:3.12.0

COPY gcloud_sdk_version.txt /tmp/gcloud_sdk_version.txt

RUN GCLOUD_SDK_VERSION=`cat /tmp/gcloud_sdk_version.txt | tr -d '\n'` && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt-get update -y && \
    apt-get install google-cloud-sdk=${GCLOUD_SDK_VERSION}-0 kubectl -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

CMD gcloud
