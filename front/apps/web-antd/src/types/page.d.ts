export type PageResult<T> = {
  list: T[];
  page: number;
  pageSize: number;
  total: number;
};
export type PageParams = {
  args?: Array<any>;
  filters?: Array<string>;
  page: number;
  pageSize: number;
  sortField?: string;
  sortOrder: 'asc' | 'desc';
};
