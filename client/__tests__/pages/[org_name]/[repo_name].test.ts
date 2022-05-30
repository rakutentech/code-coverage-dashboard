import { getServerSideProps } from "../../../pages/[org_name]/[repo_name]";

describe('[repo_name].tsx > getServerSideProps', () => {
  const apiUrl = process.env.apiURL
  const dummyCoverages = [{ dummyKey: 'dummyVal' }]
  const dummyOrgName = 'dummyOrgName'
  const dummyRepoName = 'dummyRepoName'

  const mockFetch = jest.fn(() => Promise.resolve({
    json: () => Promise.resolve(dummyCoverages)
  })) as jest.Mock
  global.fetch = mockFetch

  it('Should request with query params', async () => {
    const dummyCtx = {
      query: {
        org_name: dummyOrgName,
        repo_name: dummyRepoName,
      }
    }
    const res = await getServerSideProps(dummyCtx)

    expect(mockFetch).toHaveBeenCalled()
    expect(mockFetch).toHaveBeenCalledWith(`${apiUrl}?org_name=${dummyOrgName}&repo_name=${dummyRepoName}&full=true&per_page=10000`)
    expect(res).toEqual(
      expect.objectContaining({
        props: {
          coverages: dummyCoverages
        }
      })
    );
  })
})