import { getServerSideProps } from "../../pages/index";

describe('index.tsx > getServerSideProps', () => {
  const apiUrl = process.env.apiURL
  const dummyCoverages = [{ dummyKey: 'dummyVal' }]

  const mockFetch = jest.fn(() => Promise.resolve({
    json: () => Promise.resolve(dummyCoverages)
  })) as jest.Mock
  global.fetch = mockFetch

  it('Should request with query params', async () => {
    const dummyCtx = {
      query: {}
    }
    const res = await getServerSideProps(dummyCtx)

    expect(mockFetch).toHaveBeenCalled()
    expect(mockFetch).toHaveBeenCalledWith(`${apiUrl}?per_page=10000`)
    expect(res).toEqual(
      expect.objectContaining({
        props: {
          coverages: dummyCoverages
        }
      })
    );
  })
})