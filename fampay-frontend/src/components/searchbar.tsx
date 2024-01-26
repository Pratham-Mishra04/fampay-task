import { MagnifyingGlass } from '@phosphor-icons/react';
import { useRouter } from 'next/router';
import React, { Dispatch, SetStateAction } from 'react';

interface Props {
  search: string;
  setSearch: Dispatch<SetStateAction<string>>;
}

const SearchBar = ({ search, setSearch }: Props) => {
  const router = useRouter();

  const handleChange = (el: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(el.target.value);
  };

  const handleSubmit = (el: React.FormEvent<HTMLFormElement>) => {
    el.preventDefault();
    if (search === '') router.push('/');
    else {
      router.push(`/?message=${search}`);
    }
  };

  return (
    <div className="relative w-1/2 max-md:w-2/3">
      <form
        onSubmit={handleSubmit}
        className={`w-full h-12 px-4 text-primary_black bg-gray-100 flex items-center justify-between gap-8 mx-auto rounded-md border-gray-200 border-2 transition-ease-200`}
      >
        <input
          className="h-full grow bg-transparent focus:outline-none font-primary font-medium"
          type="text"
          placeholder="Search"
          value={search}
          onChange={handleChange}
        />
        <MagnifyingGlass size={32} className="opacity-75" />
      </form>
    </div>
  );
};

export default SearchBar;
