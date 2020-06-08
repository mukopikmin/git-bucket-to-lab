new Vue({
  el: "#app",
  data: {
    gitbucketToken: "",
    gitlabToken: ""
  },
  mounted() {
    try {
      this.gitbucketToken = localStorage.getItem('gitbucketToken');
      this.gitlabToken = localStorage.getItem('gitlabToken');
    } catch (e) {
      localStorage.removeItem('gitbucketToken');
      localStorage.removeItem('gitlabToken');
    }

    console.log(this.gitbucketToken)
  }
})