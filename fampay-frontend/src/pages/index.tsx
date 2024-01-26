import Head from 'next/head';
import Header from '@/components/header';
import Loader from '@/components/loader';
import SearchBar from '@/components/searchbar';
import getHandler from '@/handlers/get_handler';
import { Video } from '@/types';
import Toaster from '@/utils/toaster';
import { useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { SlidersHorizontal, Trash } from '@phosphor-icons/react';
import Filters from '@/components/filters';
import { useRouter } from 'next/router';
import moment from 'moment';

const buildURL = (baseUrl: string, params: object) => {
  const queryString = Object.entries(params)
    .filter(([_, value]) => value !== null && value !== '')
    .map(([key, value]) => {
      if (key == 'start' || key == 'end') {
        const formattedTime = moment(value).format('YYYY-MM-DDTHH:mm:ss[Z]');
        return `${key}=${formattedTime}`;
      }
      return `${key}=${value}`;
    })
    .join('&');

  return queryString ? `${baseUrl}&${queryString}` : baseUrl;
};

const Home = () => {
  const [videos, setVideos] = useState<Video[]>([]);
  const [loading, setLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  const [clickedOnFilters, setClickedOnFilters] = useState(false);

  const router = useRouter();

  const [search, setSearch] = useState('');

  const fetchVideos = async (pageIndex: number) => {
    const URL = buildURL(`${process.env.NEXT_PUBLIC_BACKEND_URL}?page=${pageIndex}&limit=${10}`, router.query);
    const res = await getHandler(URL);
    if (res.statusCode == 200) {
      if (pageIndex == 1) {
        const videosData: Video[] = res.data.videos || [];
        setVideos(videosData);
        if (videosData.length == 0) setHasMore(false);
        setPage(2);
      } else {
        const addedVideos = [...videos, ...(res.data.videos || [])];
        if (addedVideos.length === videos.length) setHasMore(false);
        setVideos(addedVideos);
        setPage(prev => prev + 1);
      }
      setLoading(false);
    } else {
      if (res.data.message) Toaster.error(res.data.message, 'error_toaster');
      else Toaster.error('Internal Server Error', 'error_toaster');
    }
  };

  useEffect(() => {
    const { message } = router.query;
    if (message) setSearch(message as string);

    setHasMore(true);
    fetchVideos(1);
  }, [router.query]);

  return (
    <>
      {clickedOnFilters ? <Filters setShow={setClickedOnFilters} /> : <></>}
      <Head>
        <title>Videos | Fampay</title>
      </Head>
      <Header />
      <div className="w-full flex justify-center items-center gap-6 py-4">
        <SearchBar search={search} setSearch={setSearch} />
        <SlidersHorizontal
          onClick={() => setClickedOnFilters(true)}
          className="cursor-pointer hover:bg-gray-100 rounded-full p-2 flex-center transition-ease-300"
          size={42}
          weight="duotone"
        />
      </div>

      {loading ? (
        <Loader />
      ) : (
        <InfiniteScroll
          className="w-full max-md:w-full max-md:px-4 mx-auto flex flex-col items-center gap-2"
          dataLength={videos.length}
          next={() => fetchVideos(page)}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {videos.map(video => {
            return <div key={video.id}>{video.title}</div>;
          })}
        </InfiniteScroll>
      )}
    </>
  );
};

export default Home;
