import "@/styles/globals.css";
import '@/styles/extras.tailwind.css';
import type { AppProps } from "next/app";
import { Inter } from 'next/font/google';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const inter = Inter({
  subsets: ['latin'],
  variable: '--inter-font',
});

export default function App({ Component, pageProps }: AppProps) {
  return    <main className={`${inter.variable}`}>
  <ToastContainer />
  <Component {...pageProps} />
</main>
}
