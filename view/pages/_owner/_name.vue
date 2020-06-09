<template>
  <div>
    <div class="card mt-3 mb-3">
      <div class="card-header">Repository / Project</div>
      <div class="card-body">
        <div class="row">
          <div class="col">
            <div v-if="repo">
              <h5 class="card-title">GitBucket</h5>
              <h5>{{ repo.full_name }}</h5>

              <p v-if="repo.branches.length === 0">
                No branches
              </p>

              <ul>
                <li v-for="branch in repo.branches" :key="branch.commit.sha">
                  {{ branch.name }} ({{ branch.commit.sha }})
                </li>
              </ul>
              <p>{{ repo.description }}</p>
            </div>

            <button class="btn btn-outline-primary" @click="migrateRepo">
              Migrate
            </button>
          </div>

          <div class="col">
            <div v-if="project">
              <h5 class="card-title">GitLab</h5>
              <h5>{{ project.path_with_namespace }}</h5>
              <ul>
                <li v-for="branch in project.branches" :key="branch.commit.id">
                  {{ branch.name }} ({{ branch.commit.id }})
                </li>
              </ul>
              <p>{{ project.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="card mb-3">
      <div class="card-header">Issues</div>
      <div class="card-body">
        <div class="row">
          <div class="col">
            <div v-if="repo">
              <Issues title="GitBucket" :issues="repo.issues" />

              <button class="btn btn-outline-primary" @click="migrateIssues">
                Migrate
              </button>
            </div>
          </div>

          <div class="col">
            <div v-if="project">
              <h5 class="card-title">GitLab</h5>
              <ul>
                <li v-for="issue in project.issues" :key="issue.id">
                  <a target="_blank" :href="issue.web_url"
                    >#{{ issue.iid }} {{ issue.title }} ({{ issue.state }})</a
                  >
                  <ul>
                    <li
                      v-for="comment in issue.comments"
                      :key="comment.id"
                      class="text-truncate"
                    >
                      #{{ comment.id }} {{ comment.body }}
                    </li>
                  </ul>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="card mb-3">
      <div class="card-header">Pull Requests / Merge Requests</div>
      <div class="card-body">
        <div class="row">
          <div class="col">
            <div v-if="repo">
              <h5 class="card-title">GitBucket</h5>
              <ul>
                <li v-for="pull in repo.pulls" :key="pull.id">
                  <a target="_blank" :href="pull.html_url"
                    >#{{ pull.number }} {{ pull.title }} ({{ pull.state }})</a
                  >
                  <ul>
                    <li
                      v-for="comment in pull.comments"
                      :key="comment.id"
                      class="text-truncate"
                    >
                      #{{ comment.id }} {{ comment.body }}
                    </li>
                  </ul>
                </li>
              </ul>

              <button class="btn btn-outline-primary" @click="migratePulls">
                Migrate
              </button>
            </div>
          </div>

          <div class="col">
            <div v-if="project">
              <h5 class="card-title">GitLab</h5>
              <ul>
                <li v-for="merge in project.merges" :key="merge.id">
                  <a target="_blank" :href="merge.web_url"
                    >#{{ merge.iid }} {{ merge.title }} ({{ merge.state }})</a
                  >
                  <ul>
                    <li
                      v-for="comment in merge.comments"
                      :key="comment.id"
                      class="text-truncate"
                    >
                      #{{ comment.id }} {{ comment.body }}
                    </li>
                  </ul>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Issues from '@/components/issues'

export default {
  components: {
    Issues
  },
  data: () => {
    return {
      repo: null,
      project: null
    }
  },
  async mounted() {
    const { owner, name } = this.$nuxt.$route.params

    try {
      this.gitbucketToken = localStorage.getItem('gitbucketToken')
      this.gitlabToken = localStorage.getItem('gitlabToken')
    } catch (e) {
      localStorage.removeItem('gitbucketToken')
      localStorage.removeItem('gitlabToken')
    }

    if (!(this.gitbucketToken && this.gitlabToken)) {
      this.$router.push('/auth')
    }

    const res = await axios.get(`http://localhost:1323/${owner}/${name}`, {
      headers: {
        'X-GITBUCKET-TOKEN': this.gitbucketToken,
        'X-GITLAB-TOKEN': this.gitlabToken
      }
    })

    this.repo = res.data.Repo
    this.project = res.data.Project
  },
  methods: {
    async migrateRepo() {
      const res = await axios.post(
        `http://localhost:1323/${this.repo.owner.login}/${this.repo.name}/repo`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.project = res.data.Project
    },
    async migrateIssues() {
      const res = await axios.post(
        `http://localhost:1323/${this.repo.owner.login}/${this.repo.name}/issues`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.project = res.data.Project
    },
    async migratePulls() {
      const res = await axios.post(
        `http://localhost:1323/${this.repo.owner.login}/${this.repo.name}/pulls`,
        null,
        {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        }
      )

      this.project = res.data.Project
    }
  }
}
</script>

<style></style>
