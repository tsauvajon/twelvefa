## /!\ this file was not tested, I ran all of the commands manually

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

# set zones
gcloud config set compute/zone $GOOGLE_COMPUTE_ZONE
gcloud config set compute/region $GOOGLE_COMPUTE_REGION

# enable billing
gcloud beta billing projects link ${PROJECT_ID} \
  --billing-account ${TF_VAR_billing_account}

# create service accounts
gcloud iam service-accounts create terraform \
  --display-name "Terraform admin account"
gcloud iam service-accounts create circleci \
  --display-name "CircleCI service account"

# store credentials in a JSON file
gcloud iam service-accounts keys create ${TF_CREDS} \
  --iam-account terraform@${PROJECT_ID}.iam.gserviceaccount.com

gcloud iam service-accounts keys create ${CIRCLECI_CREDS} \
  --iam-account circleci@${PROJECT_ID}.iam.gserviceaccount.com

# add the project editor role to the service account
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member serviceAccount:terraform@${PROJECT_ID}.iam.gserviceaccount.com \
  --role roles/editor

# push images to the container registry -- also gives pull access
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member serviceAccount:circleci@${PROJECT_ID}.iam.gserviceaccount.com \
  --role roles/storage.admin

# enable APIs (gcloud services list --available)
gcloud services enable compute.googleapis.com
gcloud services enable containerregistry.googleapis.com
gcloud services enable iam.googleapis.com
gcloud services enable iamcredentials.googleapis.com
gcloud services enable container.googleapis.com # GKE
