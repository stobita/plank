import React, { useCallback } from 'react';
import styled from 'styled-components';
import { useDropzone } from 'react-dropzone';

type Props = {
  images: string[];
  onChange: (image: string[]) => void;
};

export const ImageUploader = (props: Props) => {
  const onDrop = useCallback(async (acceptedFiles) => {
    const images = await Promise.all<string>(
      acceptedFiles.map(async (file: File) => {
        return await getFileAsDataURL(file);
      }),
    );
    props.onChange(images);
  }, []);
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  const getFileAsDataURL = (file: Blob): Promise<string> => {
    return new Promise((resolve) => {
      const reader = new FileReader();
      reader.onloadend = () => {
        if (typeof reader.result === 'string') {
          resolve(reader.result);
        }
      };
      reader.readAsDataURL(file);
    });
  };

  return (
    <>
      {props.images.map((v, i) => (
        <img src={v} key={i}></img>
      ))}
      <FileUploadArea {...getRootProps()}>
        <input {...getInputProps()} />
        {isDragActive ? (
          <p>Drop the files here ...</p>
        ) : (
          <p>Drop and Drop some files here ...</p>
        )}
      </FileUploadArea>
    </>
  );
};

const FileUploadArea = styled.div`
  display: flex;
  justify-content: center;
`;
