import { FileExplorerRepository } from '../repositories/FileExplorerRepository';
import { checkIf } from '../../../utils/checkIf';

class FileExplorerController {
  constructor() {
    this.repository = new FileExplorerRepository();
  }

  /**
   * description - Fetch storages
   *
   * @return {Promise<{data: object|boolean, error: string|null, stderr: string|null}>}
   */
  async listStorages({ deviceType }) {
    checkIf(deviceType, 'string');

    return this.repository.listStorages({ deviceType });
  }

  /**
   * description - Fetch files in the path
   *
   * @param {string} deviceType
   * @param {string} filePath
   * @param {string} ignoreHidden
   * @param {string} storageId
   * @return {Promise<{data: array|null, error: string|null, stderr: string|null}>}
   */
  async listFiles({ deviceType, filePath, ignoreHidden, storageId }) {
    checkIf(deviceType, 'string');
    checkIf(filePath, 'string');
    checkIf(ignoreHidden, 'boolean');

    return this.repository.listFiles({
      deviceType,
      filePath,
      ignoreHidden,
      storageId,
    });
  }

  /**
   * description - Rename a file
   *
   * @param {string} deviceType
   * @param {string} filePath
   * @param {string} newFilename
   * @param {string} storageId
   * @return {Promise<{data: null|boolean, error: string|null, stderr: string|null}>}
   */
  async renameFile({ deviceType, filePath, newFilename, storageId }) {
    checkIf(deviceType, 'string');
    checkIf(filePath, 'string');
    checkIf(newFilename, 'string');

    return this.repository.renameFile({
      deviceType,
      filePath,
      newFilename,
      storageId,
    });
  }

  /**
   * description - Delete files
   *
   * @param {string} deviceType
   * @param {[string]} fileList
   * @param {string} storageId
   * @return {Promise<{data: null|boolean, error: string|null, stderr: string|null}>}
   */
  async deleteFiles({ deviceType, fileList, storageId }) {
    checkIf(deviceType, 'string');
    checkIf(fileList, 'array');

    return this.repository.deleteFiles({
      deviceType,
      fileList,
      storageId,
    });
  }

  /**
   * description - Create a directory
   *
   * @param {string} deviceType
   * @param {string} filePath
   * @param {string} storageId
   * @return {Promise<{data: null|boolean, error: string|null, stderr: string|null}>}
   */
  async makeDirectory({ deviceType, filePath, storageId }) {
    checkIf(deviceType, 'string');
    checkIf(filePath, 'string');

    return this.repository.makeDirectory({
      deviceType,
      filePath,
      storageId,
    });
  }
}

const fileExplorerController = new FileExplorerController();

export default fileExplorerController;
