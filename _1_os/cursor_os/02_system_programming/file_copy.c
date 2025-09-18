#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <errno.h>
#include <string.h>
#include <time.h>

// Function to copy file with progress indication
int copy_file(const char *src, const char *dst) {
    struct stat src_stat;
    if (stat(src, &src_stat) == -1) {
        perror("Failed to get source file info");
        return -1;
    }
    
    int src_fd = open(src, O_RDONLY);
    if (src_fd == -1) {
        perror("Failed to open source file");
        return -1;
    }
    
    int dst_fd = open(dst, O_WRONLY | O_CREAT | O_TRUNC, src_stat.st_mode);
    if (dst_fd == -1) {
        perror("Failed to open destination file");
        close(src_fd);
        return -1;
    }
    
    char buffer[4096];
    ssize_t bytes_read, bytes_written;
    size_t total_copied = 0;
    size_t file_size = src_stat.st_size;
    
    printf("Copying %s to %s (%zu bytes)...\n", src, dst, file_size);
    
    while ((bytes_read = read(src_fd, buffer, sizeof(buffer))) > 0) {
        bytes_written = write(dst_fd, buffer, bytes_read);
        if (bytes_written != bytes_read) {
            perror("Write error");
            close(src_fd);
            close(dst_fd);
            return -1;
        }
        
        total_copied += bytes_written;
        
        // Show progress
        if (file_size > 0) {
            int progress = (total_copied * 100) / file_size;
            printf("\rProgress: %zu/%zu bytes (%d%%)", 
                   total_copied, file_size, progress);
            fflush(stdout);
        }
    }
    
    if (bytes_read == -1) {
        perror("Read error");
        close(src_fd);
        close(dst_fd);
        return -1;
    }
    
    printf("\n");
    
    // Preserve file timestamps
    struct timespec times[2];
    times[0] = src_stat.st_atim;
    times[1] = src_stat.st_mtim;
    futimens(dst_fd, times);
    
    close(src_fd);
    close(dst_fd);
    return 0;
}

// Function to copy directory recursively
int copy_directory(const char *src, const char *dst) {
    struct stat src_stat;
    if (stat(src, &src_stat) == -1) {
        perror("Failed to get source directory info");
        return -1;
    }
    
    if (!S_ISDIR(src_stat.st_mode)) {
        fprintf(stderr, "Source is not a directory: %s\n", src);
        return -1;
    }
    
    // Create destination directory
    if (mkdir(dst, src_stat.st_mode) == -1) {
        if (errno != EEXIST) {
            perror("Failed to create destination directory");
            return -1;
        }
    }
    
    // Open source directory
    DIR *dir = opendir(src);
    if (dir == NULL) {
        perror("Failed to open source directory");
        return -1;
    }
    
    struct dirent *entry;
    while ((entry = readdir(dir)) != NULL) {
        // Skip . and ..
        if (strcmp(entry->d_name, ".") == 0 || strcmp(entry->d_name, "..") == 0) {
            continue;
        }
        
        // Build full paths
        char src_path[1024];
        char dst_path[1024];
        snprintf(src_path, sizeof(src_path), "%s/%s", src, entry->d_name);
        snprintf(dst_path, sizeof(dst_path), "%s/%s", dst, entry->d_name);
        
        // Get file info
        struct stat entry_stat;
        if (stat(src_path, &entry_stat) == -1) {
            perror("Failed to get file info");
            continue;
        }
        
        if (S_ISDIR(entry_stat.st_mode)) {
            // Recursively copy directory
            printf("Copying directory: %s\n", src_path);
            if (copy_directory(src_path, dst_path) == -1) {
                closedir(dir);
                return -1;
            }
        } else {
            // Copy file
            printf("Copying file: %s\n", src_path);
            if (copy_file(src_path, dst_path) == -1) {
                closedir(dir);
                return -1;
            }
        }
    }
    
    closedir(dir);
    return 0;
}

// Function to compare two files
int compare_files(const char *file1, const char *file2) {
    int fd1 = open(file1, O_RDONLY);
    if (fd1 == -1) {
        perror("Failed to open first file");
        return -1;
    }
    
    int fd2 = open(file2, O_RDONLY);
    if (fd2 == -1) {
        perror("Failed to open second file");
        close(fd1);
        return -1;
    }
    
    char buffer1[4096];
    char buffer2[4096];
    ssize_t bytes_read1, bytes_read2;
    
    while (1) {
        bytes_read1 = read(fd1, buffer1, sizeof(buffer1));
        bytes_read2 = read(fd2, buffer2, sizeof(buffer2));
        
        if (bytes_read1 != bytes_read2) {
            close(fd1);
            close(fd2);
            return 1; // Files are different
        }
        
        if (bytes_read1 == 0) {
            break; // End of both files
        }
        
        if (memcmp(buffer1, buffer2, bytes_read1) != 0) {
            close(fd1);
            close(fd2);
            return 1; // Files are different
        }
    }
    
    close(fd1);
    close(fd2);
    return 0; // Files are identical
}

// Function to get file size
off_t get_file_size(const char *filename) {
    struct stat file_stat;
    if (stat(filename, &file_stat) == -1) {
        return -1;
    }
    return file_stat.st_size;
}

// Function to format file size for display
void format_file_size(off_t size, char *buffer, size_t buffer_size) {
    const char *units[] = {"B", "KB", "MB", "GB", "TB"};
    int unit_index = 0;
    double file_size = (double)size;
    
    while (file_size >= 1024 && unit_index < 4) {
        file_size /= 1024;
        unit_index++;
    }
    
    snprintf(buffer, buffer_size, "%.2f %s", file_size, units[unit_index]);
}

int main(int argc, char *argv[]) {
    if (argc < 3) {
        fprintf(stderr, "Usage: %s <source> <destination> [options]\n", argv[0]);
        fprintf(stderr, "Options:\n");
        fprintf(stderr, "  -r    Copy directories recursively\n");
        fprintf(stderr, "  -v    Verify copy by comparing files\n");
        fprintf(stderr, "  -h    Show this help message\n");
        return 1;
    }
    
    int recursive = 0;
    int verify = 0;
    
    // Parse options
    for (int i = 3; i < argc; i++) {
        if (strcmp(argv[i], "-r") == 0) {
            recursive = 1;
        } else if (strcmp(argv[i], "-v") == 0) {
            verify = 1;
        } else if (strcmp(argv[i], "-h") == 0) {
            fprintf(stderr, "Usage: %s <source> <destination> [options]\n", argv[0]);
            fprintf(stderr, "Options:\n");
            fprintf(stderr, "  -r    Copy directories recursively\n");
            fprintf(stderr, "  -v    Verify copy by comparing files\n");
            fprintf(stderr, "  -h    Show this help message\n");
            return 0;
        } else {
            fprintf(stderr, "Unknown option: %s\n", argv[i]);
            return 1;
        }
    }
    
    const char *src = argv[1];
    const char *dst = argv[2];
    
    // Check if source exists
    struct stat src_stat;
    if (stat(src, &src_stat) == -1) {
        perror("Source does not exist");
        return 1;
    }
    
    // Show file information
    if (S_ISREG(src_stat.st_mode)) {
        char size_str[32];
        format_file_size(src_stat.st_size, size_str, sizeof(size_str));
        printf("Source file: %s (%s)\n", src, size_str);
    } else if (S_ISDIR(src_stat.st_mode)) {
        printf("Source directory: %s\n", src);
    }
    
    // Perform copy
    int result;
    if (S_ISDIR(src_stat.st_mode) && recursive) {
        result = copy_directory(src, dst);
    } else if (S_ISREG(src_stat.st_mode)) {
        result = copy_file(src, dst);
    } else {
        fprintf(stderr, "Source is a directory but -r option not specified\n");
        return 1;
    }
    
    if (result == -1) {
        fprintf(stderr, "Copy failed\n");
        return 1;
    }
    
    printf("Copy completed successfully\n");
    
    // Verify copy if requested
    if (verify && S_ISREG(src_stat.st_mode)) {
        printf("Verifying copy...\n");
        if (compare_files(src, dst) == 0) {
            printf("Files are identical - copy verified\n");
        } else {
            printf("Files are different - copy verification failed\n");
            return 1;
        }
    }
    
    return 0;
}
