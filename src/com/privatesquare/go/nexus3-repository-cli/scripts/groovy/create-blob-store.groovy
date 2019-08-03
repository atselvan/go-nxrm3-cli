// create-blob-store.groovy is a  Nexus3 Integration API definition to create a blob store in Nexus
// This definition currently only supports file system blob stores

// TODO : Add support for adding S3 blob stores

// import libraries for json parsing
import groovy.json.JsonOutput
import groovy.json.JsonSlurper

// input map
blob = new JsonSlurper().parseText(args)

existingBlobStore = blobStore.getBlobStoreManager().get(blob.name)
if (existingBlobStore == null) {
    blobStore.createFileBlobStore(blob.name, blob.path)
}

// output map
Map output = [:]
// output request status
output.put('status', '200 OK')
// output blob details
output.put('name', blob.name)
output.put('url', blob.url)

JsonOutput.toJson(repodata)