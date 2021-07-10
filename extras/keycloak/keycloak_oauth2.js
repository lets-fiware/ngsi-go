function introspectAccessToken(req) {
    req.subrequest("/_request_introspection",
        function(reply) {
            req.return(reply.status == 200
                ? ((JSON.parse(reply.responseBody).active == true) ? 204 : 403)
                : 401);
        }
    );
}
