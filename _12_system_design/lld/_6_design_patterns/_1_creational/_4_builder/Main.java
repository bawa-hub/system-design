package _6_design_patterns._1_creational._4_builder;

import java.util.HashMap;
import java.util.Map;

// Step 1: Define the Product (HTTP Request)

// The object being constructed.
class HttpRequest {
    private String url;
    private String method;
    private Map<String, String> headers;
    private String body;

    // Private constructor to enforce immutability
    private HttpRequest(Builder builder) {
        this.url = builder.url;
        this.method = builder.method;
        this.headers = builder.headers;
        this.body = builder.body;
    }

    // Nested static Builder class
    public static class Builder {
        private String url;
        private String method;
        private Map<String, String> headers = new HashMap<>();
        private String body;

        public Builder(String url) { // Mandatory field
            this.url = url;
        }

        public Builder method(String method) {
            this.method = method;
            return this;
        }

        public Builder addHeader(String key, String value) {
            this.headers.put(key, value);
            return this;
        }

        public Builder body(String body) {
            this.body = body;
            return this;
        }

        public HttpRequest build() {
            return new HttpRequest(this);
        }
    }

    @Override
    public String toString() {
        return "HttpRequest{" +
                "url='" + url + '\'' +
                ", method='" + method + '\'' +
                ", headers=" + headers +
                ", body='" + body + '\'' +
                '}';
    }
}

// Step 2: Use the Builder to Create the HTTP Request

// Construct the object step by step.
public class Main {
    public static void main(String[] args) {
        HttpRequest request = new HttpRequest.Builder("https://example.com")
                .method("POST")
                .addHeader("Content-Type", "application/json")
                .addHeader("Authorization", "Bearer token123")
                .body("{\"key\": \"value\"}")
                .build();

        System.out.println(request);
    }
}

// Output:
// HttpRequest{
//     url='https://example.com',
//     method='POST',
//     headers={Content-Type=application/json, Authorization=Bearer token123},
//     body='{"key": "value"}'
// }

