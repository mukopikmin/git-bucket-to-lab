<template>
  <div>
    <b-row>
      <b-col offset-sm="2" sm="8">
        <p>
          Set username of GitBucket and GitLab. This application can be used
          when same user exists on both services. If there is no user which has
          same login name, set up before start migration.
        </p>
        <b-input-group prepend="Username" class="mt-3">
          <b-form-input v-model="usernameInput"></b-form-input>
        </b-input-group>

        <div v-if="isUserSet">
          <p class="mt-5">
            Generate personal access token for GitBucket and GitLab. You need to
            generate token for GitLab with API permission.
          </p>

          <div v-if="loading" class="text-center">
            <b-spinner variant="primary"></b-spinner>
          </div>

          <div v-else>
            <b-input-group prepend="GitBucket" class="mt-3">
              <b-form-input v-model="gitbucketTokenInput"></b-form-input>
              <b-input-group-append>
                <b-button
                  variant="outline-primary"
                  @click="generateGitbucketToken"
                >
                  <b-icon-box-arrow-up-right
                    class="mr-1"
                  ></b-icon-box-arrow-up-right>
                  Generate
                </b-button>
              </b-input-group-append>
            </b-input-group>

            <b-input-group prepend="GitLab" class="mt-3">
              <b-form-input v-model="gitlabTokenInput"></b-form-input>
              <b-input-group-append>
                <b-button
                  variant="outline-primary"
                  @click="generateGitlabToken"
                >
                  <b-icon-box-arrow-up-right
                    class="mr-1"
                  ></b-icon-box-arrow-up-right>
                  Generate
                </b-button>
              </b-input-group-append>
            </b-input-group>

            <b-button
              class="my-3"
              block
              variant="outline-primary"
              @click="auth"
            >
              Authorize
            </b-button>
          </div>
        </div>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
  props: {
    loading: Boolean
  },
  data() {
    return {
      usernameInput: '',
      gitbucketTokenInput: '',
      gitlabTokenInput: ''
    }
  },
  computed: {
    ...mapState([
      'gitbucketUser',
      'gitbucketUrl',
      'gitlabUrl',
      'gitbucketToken',
      'gitlabToken',
      'username'
    ]),
    isUserSet() {
      return this.usernameInput !== ''
    }
  },
  watch: {
    gitbucketUser: {
      immediate: true,
      handler(pre, cur) {
        setTimeout(() => {
          this.usernameInput = this.username
        }, 0)
      }
    },
    gitbucketToken: {
      immediate: true,
      handler() {
        this.gitbucketTokenInput = this.gitbucketToken
      }
    },
    gitlabToken: {
      immediate: true,
      handler() {
        this.gitlabTokenInput = this.gitlabToken
      }
    }
  },
  methods: {
    ...mapActions(['setGitbucketToken', 'setGitlabToken']),
    auth() {
      this.setGitbucketToken(this.gitbucketTokenInput)
      this.setGitlabToken(this.gitlabTokenInput)

      this.$router.push('/')
    },
    generateGitbucketToken() {
      window.open(
        `${this.gitbucketUrl}/${this.username}/_application`,
        '_blank'
      )
    },
    generateGitlabToken() {
      window.open(`${this.gitlabUrl}/profile/personal_access_tokens`, '_blank')
    }
  }
}
</script>

<style scoped>
.input-group-text {
  min-width: 6em;
}
</style>
