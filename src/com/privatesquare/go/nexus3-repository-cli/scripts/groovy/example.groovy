import groovy.json.JsonOutput

def repo = repository.createMavenHosted("test21")

def repodata = [:]
repodata.put("name", repo.name)
repodata.put("url", repo.url)

return JsonOutput.toJson(repodata)


repository.createMavenHosted('maven-internal')


import org.sonatype.nexus.blobstore.api.BlobStoreManager
import org.sonatype.nexus.repository.*

repository.createMavenHosted('private')



blobStore.createFileBlobStore('npm', 'npm')

security.addRole('blue','blue', 'Blue Role', [], [])

repository.createBowerHosted('bower-internal')

core.baseUrl('http://repo.example.com')