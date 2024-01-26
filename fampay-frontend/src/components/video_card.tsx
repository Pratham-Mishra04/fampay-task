import { Video } from '@/types';
import React from 'react';
import Image from 'next/image';
import moment from 'moment';
import Link from 'next/link';

interface Props {
  video: Video;
}

const VideoCard = ({ video }: Props) => {
  return (
    <Link
      href={`https://youtube.com/${video.youtubeID}`}
      target="_blank"
      className="w-72 h-48 max-lg:w-60 max-lg:h-60 max-md:w-72 max-md:h-72 rounded-lg relative group transition-ease-out-500"
    >
      <div className="w-full h-full  absolute top-0 hidden group-hover:flex animate-fade_third justify-end z-[6] rounded-lg p-1">
        {/* <BookmarkSimple size={24}  className="opacity-75" /> */}
      </div>
      <div className="w-full h-full rounded-lg overflow-clip p-4 text-sm backdrop-blur-xl text-white absolute top-0 left-0 bg-gradient-to-b from-[#00000080] z-[5] to-transparent opacity-0 group-hover:opacity-100 transition-ease-300"></div>
      <div className="w-full h-full rounded-lg overflow-clip p-4 text-sm fade-img backdrop-blur-sm text-white absolute top-0 left-0 z-[5] opacity-0 group-hover:opacity-100 transition-ease-300">
        {video.description}
      </div>
      <Image
        crossOrigin="anonymous"
        className="w-full h-full rounded-lg object-cover absolute top-0 left-0 "
        src={video.thumbnail}
        alt="Video Thumbnail"
        width={500}
        height={500}
      />
      <div className="w-full glassMorphism text-white rounded-b-lg font-primary absolute bottom-0 right-0 flex flex-col gap-2 px-4 py-2">
        <div className="line-clamp-1 text-lg font-medium">{video.title}</div>
        <div className="w-full flex items-center justify-between">
          <div className="w-fit flex items-center gap-1 line-clamp-1 text-sm">{video.channelTitle}</div>
          <div className="text-xs">{moment(video.uploadedAt).fromNow()}</div>
        </div>
      </div>
    </Link>
  );
};

export default VideoCard;
