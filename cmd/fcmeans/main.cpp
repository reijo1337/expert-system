#include "fcm.h"
#include <cstdlib>
#include <vector>
#include <iostream>
#include <iomanip>

constexpr int ClusterCount = 3;
constexpr int ColumnCount = 3;

std::vector<std::vector<double>> dataset = {
    { 5.1, 3.5, 0.2 },
    { 4.9, 3., 0.2 },
    { 4.7, 3.2, 0.2 },
    { 4.6, 3.1, 0.2 },
    { 5., 3.6, 0.2 },
    { 5.4, 3.9, 0.4 },
    { 4.6, 3.4, 0.3 },
    { 5., 3.4, 0.2 },
    { 4.4, 2.9, 0.2 },
    { 4.9, 3.1, 0.1 },
    { 5.4, 3.7, 0.2 },
    { 4.8, 3.4, 0.2 },
    { 4.8, 3., 0.1 },
    { 4.3, 3., 0.1 },
    { 5.8, 4., 0.2 },
    { 5.7, 4.4, 0.4 },
    { 5.4, 3.9, 0.4 },
    { 5.1, 3.5, 0.3 },
    { 5.7, 3.8, 0.3 },
    { 5.1, 3.8, 0.3 },
    { 5.4, 3.4, 0.2 },
    { 5.1, 3.7, 0.4 },
    { 4.6, 3.6, 0.2 },
    { 5.1, 3.3, 0.5 },
    { 4.8, 3.4, 0.2 },
    { 5., 3., 0.2 },
    { 5., 3.4, 0.4 },
    { 5.2, 3.5, 0.2 },
    { 5.2, 3.4, 0.2 },
    { 4.7, 3.2, 0.2 },
    { 4.8, 3.1, 0.2 },
    { 5.4, 3.4, 0.4 },
    { 5.2, 4.1, 0.1 },
    { 5.5, 4.2, 0.2 },
    { 4.9, 3.1, 0.2 },
    { 5., 3.2, 0.2 },
    { 5.5, 3.5, 0.2 },
    { 4.9, 3.6, 0.1 },
    { 4.4, 3., 0.2 },
    { 5.1, 3.4, 0.2 },
    { 5., 3.5, 0.3 },
    { 4.5, 2.3, 0.3 },
    { 4.4, 3.2, 0.2 },
    { 5., 3.5, 0.6 },
    { 5.1, 3.8, 0.4 },
    { 4.8, 3., 0.3 },
    { 5.1, 3.8, 0.2 },
    { 4.6, 3.2, 0.2 },
    { 5.3, 3.7, 0.2 },
    { 5., 3.3, 0.2 },
    { 7., 3.2, 1.4 },
    { 6.4, 3.2, 1.5 },
    { 6.9, 3.1, 1.5 },
    { 5.5, 2.3, 1.3 },
    { 6.5, 2.8, 1.5 },
    { 5.7, 2.8, 1.3 },
    { 6.3, 3.3, 1.6 },
    { 4.9, 2.4, 1. },
    { 6.6, 2.9, 1.3 },
    { 5.2, 2.7, 1.4 },
    { 5., 2., 1. },
    { 5.9, 3., 1.5 },
    { 6., 2.2, 1. },
    { 6.1, 2.9, 1.4 },
    { 5.6, 2.9, 1.3 },
    { 6.7, 3.1, 1.4 },
    { 5.6, 3., 1.5 },
    { 5.8, 2.7, 1. },
    { 6.2, 2.2, 1.5 },
    { 5.6, 2.5, 1.1 },
    { 5.9, 3.2, 1.8 },
    { 6.1, 2.8, 1.3 },
    { 6.3, 2.5, 1.5 },
    { 6.1, 2.8, 1.2 },
    { 6.4, 2.9, 1.3 },
    { 6.6, 3., 1.4 },
    { 6.8, 2.8, 1.4 },
    { 6.7, 3., 1.7 },
    { 6., 2.9, 1.5 },
    { 5.7, 2.6, 1. },
    { 5.5, 2.4, 1.1 },
    { 5.5, 2.4, 1. },
    { 5.8, 2.7, 1.2 },
    { 6., 2.7, 1.6 },
    { 5.4, 3., 1.5 },
    { 6., 3.4, 1.6 },
    { 6.7, 3.1, 1.5 },
    { 6.3, 2.3, 1.3 },
    { 5.6, 3., 1.3 },
    { 5.5, 2.5, 1.3 },
    { 5.5, 2.6, 1.2 },
    { 6.1, 3., 1.4 },
    { 5.8, 2.6, 1.2 },
    { 5., 2.3, 1. },
    { 5.6, 2.7, 1.3 },
    { 5.7, 3., 1.2 },
    { 5.7, 2.9, 1.3 },
    { 6.2, 2.9, 1.3 },
    { 5.1, 2.5, 1.1 },
    { 5.7, 2.8, 1.3 },
    { 6.3, 3.3, 2.5 },
    { 5.8, 2.7, 1.9 },
    { 7.1, 3., 2.1 },
    { 6.3, 2.9, 1.8 },
    { 6.5, 3., 2.2 },
    { 7.6, 3., 2.1 },
    { 4.9, 2.5, 1.7 },
    { 7.3, 2.9, 1.8 },
    { 6.7, 2.5, 1.8 },
    { 7.2, 3.6, 2.5 },
    { 6.5, 3.2, 2. },
    { 6.4, 2.7, 1.9 },
    { 6.8, 3., 2.1 },
    { 5.7, 2.5, 2. },
    { 5.8, 2.8, 2.4 },
    { 6.4, 3.2, 2.3 },
    { 6.5, 3., 1.8 },
    { 7.7, 3.8, 2.2 },
    { 7.7, 2.6, 2.3 },
    { 6., 2.2, 1.5 },
    { 6.9, 3.2, 2.3 },
    { 5.6, 2.8, 2. },
    { 7.7, 2.8, 2. },
    { 6.3, 2.7, 1.8 },
    { 6.7, 3.3, 2.1 },
    { 7.2, 3.2, 1.8 },
    { 6.2, 2.8, 1.8 },
    { 6.1, 3., 1.8 },
    { 6.4, 2.8, 2.1 },
    { 7.2, 3., 1.6 },
    { 7.4, 2.8, 1.9 },
    { 7.9, 3.8, 2. },
    { 6.4, 2.8, 2.2 },
    { 6.3, 2.8, 1.5 },
    { 6.1, 2.6, 1.4 },
    { 7.7, 3., 2.3 },
    { 6.3, 3.4, 2.4 },
    { 6.4, 3.1, 1.8 },
    { 6., 3., 1.8 },
    { 6.9, 3.1, 2.1 },
    { 6.7, 3.1, 2.4 },
    { 6.9, 3.1, 2.3 },
    { 5.8, 2.7, 1.9 },
    { 6.8, 3.2, 2.3 },
    { 6.7, 3.3, 2.5 },
    { 6.7, 3., 2.3 },
    { 6.3, 2.5, 1.9 },
    { 6.5, 3., 2. },
    { 6.2, 3.4, 2.3 },
    { 5.9, 3., 1.8 }
};

void Draw(std::vector<std::vector<double>> dataset, std::unique_ptr<FCM> fcm)
{
    std::clog << "+===+==========+==========+==========+" << std::endl;
    std::clog << "| > | cluster0 | cluster1 | cluster2 |" << std::endl;
    std::clog << "+===+==========+==========+==========+" << std::endl;
    for (size_t i = 0; i < dataset.size(); ++i) {
        const auto& point = dataset[i];

        int maxIndex = 0;
        double cachedMaxValue = 0;
        for (size_t j = 0; j < ClusterCount; ++j) {
            if ((*fcm->get_membership())(i, j) > cachedMaxValue) {
                maxIndex = j;
                cachedMaxValue = (*fcm->get_membership())(i, j);
            }
        }

        std::clog.setf(std::ios::fixed, std::ios::floatfield);
        std::clog << "| " << maxIndex << " |";
        for (size_t j = 0; j < ClusterCount; ++j) {
            std::clog << " " << std::setprecision(6) << (*fcm->get_membership())(i, j) << " " << "|";

        }
        std::clog << std::endl << "+---+----------+----------+----------+" << std::endl;
        std::cout << point[0] << " " << point[1] << " " << point[2] << " " << maxIndex << '\n';
    }
}

int main(int argc, char** argv)
{
    constexpr double Eps = 0.5;

    auto fcm = std::make_unique<FCM>(2, Eps);
    auto* dataMatrix = new MatrixXf();
    dataMatrix->resize(dataset.size(), ColumnCount);

    for (size_t i = 0; i < dataset.size(); ++i) {
        for (size_t j = 0; j < dataset[i].size(); ++j) {
            (*dataMatrix)(i, j) = (float)dataset[i][j];
        }
    }

    fcm->set_data(dataMatrix);
    fcm->set_num_clusters(ClusterCount);

    auto* membership = new MatrixXf;
    membership->resize(dataMatrix->rows(), ClusterCount);
    for (int i = 0; i < dataMatrix->rows(); ++i) {
        std::vector<int> vector1(ClusterCount, 0);
        vector1[rand() % (vector1.size())] = 1;
        for (int j = 0; j < vector1.size(); ++j) {
            (*membership)(i, j) = vector1[j];
        }
    }
    fcm->set_membership(membership);
    fcm->compute_centers();

    int maxIterations = 100;
    while (maxIterations--) {
        auto accuracy = fcm->update_membership();
        std::clog << "Accuracy " << accuracy << (accuracy <= Eps ? " <= " : " > ") << Eps << " eps" << std::endl;
        if (accuracy <= Eps) {
            break;
        }
    }

    Draw(dataset, std::move(fcm));


    return 0;
}
