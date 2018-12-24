window.onload = function () {
    var webScraperApp = new Vue({
        el: '#DivMain',
        data: {
            metaData: {},
            loading: false
        },
        methods: {
            scrapePage: function () {
                var errorProtocol = function (error) {
                    console.error(error);
                    Materialize.toast('An Error Occured!', 4000);
                    this.loading = false;
                };
                try {
                    var endpoint = '/v1/topics?offset=0&limit=20';
                    this.loading = true;
                    this.$http.get(endpoint)
                        .then(function (response) {
                            if (response.status == 200) {
                                this.loading = false;
                                this.topics = JSON.parse(response.bodyText);
                                if (this.topics.length == 0) {
                                    console.log(this.topics);
                                    this.noContent = true;
                                } else {
                                    this.noContent = false;
                                }
                            } else {
                                errorProtocol(response.statusText);
                            }
                        }, function (error) {
                            errorProtocol(error)
                        });
                } catch (exp) {
                    errorProtocol(exp);
                }
            },

        }
    });
};