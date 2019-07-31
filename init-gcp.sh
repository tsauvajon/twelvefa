install_sdk()
{
    INSTALL_DIR=/home/$USER/gcp

    mkdir -p $INSTALL_DIR
    cd $INSTALL_DIR

    VERSION=google-cloud-sdk-256.0.0-linux-x86_64
    wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/${VERSION}.tar.gz

    tar -xvf ${VERSION}.tar.gz

    cd google-cloud-sdk
    ./install.sh
}

# install the GCP sdk if not present
if [ $(command -v gcloud) == "" ];then
    install_sdk
    echo "please restart the terminal to apply changes"
    exit 0
fi

gcloud auth login

# create project
gcloud projects create ${PROJECT_ID} --set-as-default
gcloud config set project ${PROJECT_ID}

# enable billing
gcloud beta billing projects link ${PROJECT_ID} \
  --billing-account ${TF_VAR_billing_account}

# create service account
gcloud iam service-accounts create terraform \
  --display-name "Terraform admin account"

# store credentials in a JSON file
gcloud iam service-accounts keys create ${TF_CREDS} \
  --iam-account terraform@${PROJECT_ID}.iam.gserviceaccount.com

# add roles
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member serviceAccount:terraform@${PROJECT_ID}.iam.gserviceaccount.com \
  --role roles/viewer

gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member serviceAccount:terraform@${PROJECT_ID}.iam.gserviceaccount.com \
  --role roles/storage.admin

# enable APIs
gcloud services enable compute.googleapis.com

