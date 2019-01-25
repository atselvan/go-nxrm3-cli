import groovy.json.JsonOutput
import groovy.json.JsonSlurper

blob = new JsonSlurper().parseText(args)

existingBlobStore = blobStore.getBlobStoreManager().get(blob.name)
if (existingBlobStore == null) {
    blobStore.createFileBlobStore(blob.name, blob.path)
}

def blobdata = [:]
blobdata.put("name", blob.name)
blobdata.put("url", blob.url)

return JsonOutput.toJson(repodata)