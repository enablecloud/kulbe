# kulbe
Application Composition


#Compilation
Need Glide to manage dependencies
To generate DeepCopy :

    glide update --force ;glide install
    hack/update-codegen.sh

To regular compilation

    glide update --force ;glide install --strip-vendor;go build
