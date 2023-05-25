import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';
import { current } from '@reduxjs/toolkit';

export const plantApi = createApi({
  reducerPath: 'plantApi',
  refetchOnReconnect: true,
  refetchOnFocus: true,
  baseQuery: fetchBaseQuery({
    baseUrl: 'http://localhost:9000/api/v1',
  }),
  endpoints: (builder) => ({
    getPlants: builder.query({
      query: () => ({
        url: 'plants',
      }),
    }),
  }),
});

export const {
  useGetPlantsQuery,
} = plantApi;
