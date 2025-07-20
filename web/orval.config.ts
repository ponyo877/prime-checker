import { defineConfig } from 'orval';

export default defineConfig({
  api: {
    input: '../typespec/tsp-output/@typespec/openapi3/openapi.yaml',
    output: {
      target: './src/generated-client/primeApi.ts',
      client: 'react-query',
      mode: 'split',
      baseUrl: 'http://localhost:8080',
    },
  },
});