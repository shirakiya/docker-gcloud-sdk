FROM python:3.13.5

COPY gcloud_sdk_version.txt /tmp/gcloud_sdk_version.txt

RUN GCLOUD_SDK_VERSION=`cat /tmp/gcloud_sdk_version.txt | tr -d '\n'` && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg && \
    apt-get update -y && \
    apt-get install -y google-cloud-cli=${GCLOUD_SDK_VERSION}-0 kubectl && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

CMD gcloud
