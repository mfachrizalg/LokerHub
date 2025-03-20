"use client"

import React, {useState} from "react"
import Link from "next/link"
import Image from 'next/image';
import { Button } from '@/components/ui/button';
import { GoX } from "react-icons/go";
import { useRouter } from "next/navigation"


export default function CompleteProfileKandidat() {
    const router = useRouter()

    const [selectedProfilePic, setSelectedProfilePic] = useState<File | null>(null);
    const [selectedCV, setSelectedCV] = useState<File | null>(null);

    const handleProfilePicChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedProfilePic(event.target.files[0]);
        }
    };

    const handleCVChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.files && event.target.files.length > 0) {
            setSelectedCV(event.target.files[0]);
        }
    };

    const handleRemoveProfilePic = () => {
        setSelectedProfilePic(null);
    };

    const handleRemoveCV = () => {
        setSelectedCV(null);
    };

    return(
        <div className="min-h-screen bg-[#F5EFE7] py-8">
            <div className="container mx-auto px-4 max-w-xl">
                <div className="my-8 flex justify-center">
                    <Image
                        src="/LokerHub_Logo.svg"
                        alt="logo"
                        className="object-contain"
                        width={400}
                        height={100}
                    />
                </div>

                <h1 className="text-3xl font-bold text-center text-[#1e3a5f] mb-8">Lengkapi Data Diri</h1>

                <form>
                <div className="bg-[#d9d9d9] bg-opacity-30 rounded-lg p-6 mb-6 border border-blue-200">
                    <h2 className="text-xl font-bold text-[#1e3a5f] mb-4">Informasi Pribadi</h2>

                    <div className="space-y-4 pl-3">
                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Nama Lengkap</label>
                        <input
                        type="text"
                        name="namaLengkap"
                        required
                        className="w-full px-3 py-2 bg-white border border-white rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                        />
                    </div>

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Deskripsi Diri</label>
                        <input
                        type="text"
                        name="deskripsiDiri"
                        required
                        className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                        />
                    </div>

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Nomor Telepon</label>
                        <input
                        type="telepon"
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

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Pendidikan Terakhir</label>
                        <input
                        type="text"
                        name="pendidikanTerakhir"
                        required
                        className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                        />
                    </div>

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Bidang yang Diminati</label>
                        <input
                        type="text"
                        name="bidangDiminati"
                        required
                        className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                        />
                    </div>

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">Lokasi</label>
                        <input
                        type="text"
                        name="lokasi"
                        required
                        className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-[#1e3a5f]"
                        />
                    </div>

                    <div>
                        <label className="block font-medium text-gray-700 mb-1 pl-4">CV</label>
                        <div className="w-full px-3 py-2 bg-white border border-gray-300 rounded-md focus:ring-2 focus:ring-[#1e3a5f]">
                            {selectedCV ? (
                            <div className="inline-flex items-center bg-blue-200 px-3 py-1 rounded-md text-sm border border-black">
                                <span className="mr-2">{selectedCV.name}</span>
                                <button onClick={handleRemoveCV} className="text-gray-600 hover:text-gray-800">
                                <GoX className="w-4 h-4" />
                                </button>
                            </div>
                            ) : (
                                <label htmlFor="cv-upload" className="cursor-pointer text-gray-500">
                                    Pilih file...
                                </label>
                            )}
                            <input
                            id="cv-upload"
                            type="file"
                            accept=".pdf"
                            className="hidden"
                            onChange={handleCVChange}
                            />
                        </div>
                    </div>
                    </div>
                </div>
                </form>

                <div className="flex justify-center">
                    <Button type="submit"
                            className="py-2 px-8 bg-[#a8c5e2] border border-black hover:bg-[#97b8d9] text-[#1e3a5f] font-medium rounded-md transition-colors">
                        Submit
                    </Button>
                </div>
            </div>
        </div>

    )
}