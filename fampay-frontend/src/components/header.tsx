import Image from 'next/image';
import React from 'react';

const Header = () => {
  return (
    <div className="w-full bg-white p-4 flex items-center border-b-[1px] border-dashed border-primary_black">
      <Image src="/fampay-logo.png" alt="" width={100} height={100} className="w-12 h-fit" />
      <div className="h-fit font-primary font-medium text-3xl">videos</div>
    </div>
  );
};

export default Header;
