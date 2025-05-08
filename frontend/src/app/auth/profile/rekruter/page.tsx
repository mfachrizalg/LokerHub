"use client"

import React, {useState} from "react"
import Link from "next/link"
import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { GoX } from "react-icons/go";
import { useRouter } from "next/navigation"
import { ArrowLeft } from "lucide-react"
import Footer from "@/components/global/footer";


export default function CompleteProfileRekruter() {
    const router = useRouter()
    const [showForm, setShowForm] = useState(false);
    const [selectedProfilePic, setSelectedProfilePic] = useState<File | null>(null);
    const [selectedCompanyLogo, setSelectedCompanyLogo] = useState<File | null>(null);

     // Function to upload profile photos
     const handleProfilePicChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedProfilePic(event.target.files[0]);
        }
    };

    // Function to upload company logo
    const handleCompanyLogoChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedCompanyLogo(event.target.files[0]);
        }
    };

    // Function to delete profile photo
    const handleRemoveProfilePic = () => {
        setSelectedProfilePic(null);
    };

    // Function to remove company logo
    const handleRemoveCompanyLogo = () => {
        setSelectedCompanyLogo(null);
    };

    return(
        <div className="min-h-screen bg-white pt-8">
            {/* Back button */}
            <div className="ml-20">
                <button onClick={() => router.back()} className="flex items-center text-[#1e3a5f] hover:underline cursor-pointer">
                    <ArrowLeft className="h-4 w-4 mr-1" />
                    Kembali
                </button>
            </div>
            <div className="container mx-auto px-4 max-w-xl">
                <div className="my-8 flex justify-center">
                    <Image
                        src="/LokerHub_Logo.svg"
                        alt="logo"
                        className="object-contain"
                        width={320}
                        height={80}
                    />
                </div>

                <h1 className="text-2xl font-bold text-center text-[#1e3a5f] mb-3">Lengkapi Data Diri</h1>

                {/* Form for personal information */}
                <form>
                <div className="bg-[#f8f0e5] bg-opacity-30 rounded-lg p-6 mb-6 border border-[#e8d5c4]">
                    <h2 className="text-xl font-bold text-[#1e3a5f] mb-4">Informasi Pribadi</h2>

                    <div className="space-y-4 pl-3">
                        <div>
                            <label className="block font-medium text-gray-700 mb-1 pl-4">Nama Lengkap</label>
                            <input
                            type="text"
                            name="namaLengkap"
                            required
                            className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <div>
                            <label className="block font-medium text-gray-700 mb-1 pl-4">Jabatan</label>
                            <input
                            type="text"
                            name="jabatan"
                            required
                            className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <div>
                            <label className="block font-medium text-gray-700 mb-1 pl-4">Nomor Telepon</label>
                            <input
                            type="text"
                            name="nomorTelepon"
                            required
                            className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                            />
                        </div>

                        <div>
                            <label className="block font-medium text-gray-700 mb-1 pl-4">Unggah Foto Profil</label>
                            <div className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:ring-2 focus:ring-[#1e3a5f]">
                                {selectedProfilePic ? (
                                <div className="inline-flex items-center bg-blue-200 px-3 py-1 rounded-md text-sm border border-black">
                                    <span className="mr-2">{selectedProfilePic.name}</span>
                                    <button onClick={handleRemoveProfilePic} className="text-gray-600 hover:text-gray-800">
                                    <GoX className="w-4 h-4" />
                                    </button>
                                </div>
                                ) : (
                                    <label htmlFor="file-upload" className="cursor-pointer text-gray-500">
                                        Pilih file...
                                    </label>
                                )}
                                <input
                                id="file-upload"
                                type="file"
                                accept="image/png, image/jpg, image/jpeg"
                                className="hidden"
                                onChange={handleProfilePicChange}
                                />
                            </div>
                        </div>

                        {!showForm && (
                            <div>
                                <div>
                                    <label className="block font-medium text-gray-700 mb-1 pl-4">Nama Perusahaan</label>
                                    <input
                                        type="text"
                                        name="namaPerusahaan"
                                        required
                                        className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                                    />
                                </div>

                                <div className="text-center flex justify-center">
                                <p className="px-4">Tidak menemukan perusahaan Anda?</p>
                                <button
                                    onClick={() => setShowForm(!showForm)}
                                    className="font-medium hover:underline cursor-pointer"
                                >
                                    Daftarkan sekarang
                                </button>
                                </div>
                            </div>
                        )}
                    </div>
                </div>
                </form>

                {/* Form for company information */}
                {showForm && (
                    <form>
                        <div className="bg-[#f8f0e5] rounded-lg p-6 mb-6 border border-[#e8d5c4] relative">
                            <button 
                                className="absolute top-3 right-3 text-gray-600 hover:text-gray-800 cursor-pointer"
                                onClick={() => setShowForm(false)}
                            >
                                <GoX className="w-6 h-6" />
                            </button>
                            <h2 className="text-xl font-bold text-[#1e3a5f] mb-4">Informasi Perusahaan</h2>
                            <div className="space-y-4 pl-3">
                                <div>
                                    <label className="block font-medium text-gray-700 mb-1 pl-4">Nama Perusahaan</label>
                                    <input 
                                    type="text" 
                                    name="namaPerusahaan" 
                                    required 
                                    className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]" />
                                </div>
                                <div>
                                    <label className="block font-medium text-gray-700 mb-1 pl-4">Unggah Logo Perusahaan</label>
                                    <div className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:ring-2 focus:ring-[#1e3a5f]">
                                        {selectedCompanyLogo ? (
                                        <div className="inline-flex items-center bg-blue-200 px-3 py-1 rounded-md text-sm border border-black">
                                            <span className="mr-2">{selectedCompanyLogo.name}</span>
                                            <button onClick={handleRemoveCompanyLogo} className="text-gray-600 hover:text-gray-800">
                                            <GoX className="w-4 h-4" />
                                            </button>
                                        </div>
                                        ) : (
                                            <label htmlFor="file-upload" className="cursor-pointer text-gray-500">
                                                Pilih file...
                                            </label>
                                        )}
                                        <input
                                        id="file-upload2"
                                        type="file"
                                        accept="image/png, image/jpg, image/jpeg"
                                        className="hidden"
                                        onChange={handleCompanyLogoChange}
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block font-medium text-gray-700 mb-1 pl-4">Industri Perusahaan</label>
                                    <input 
                                    type="text" 
                                    name="industriPerusahaan" 
                                    required 
                                    className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]" />
                                </div>
                                <div>
                                    <label className="block font-medium text-gray-700 mb-1 pl-4">Lokasi Perusahaan</label>
                                    <input 
                                    type="text" 
                                    name="lokasi" 
                                    required 
                                    className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]" />
                                </div>
                            </div>
                        </div>
                    </form>
                )}

                <div className="flex justify-center mt-6">
                    <Button type="submit"
                            className="py-2 px-6 bg-[#1e3a5f] hover:bg-[#1e3a5f]/90 text-white font-medium rounded-md transition-colors cursor-pointer mb-10">
                        Submit
                    </Button>
                </div>
            </div>
            <Footer />
        </div>

    )
}